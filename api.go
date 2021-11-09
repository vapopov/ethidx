package ethidx

import (
	"fmt"
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gofiber/fiber/v2"
)

type httpError struct {
	status  string
}

// Api base api handlers.
// @title Test api
// @version 1.0
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email me@vpopov.org
// @host localhost:3000
// @BasePath /
type Api struct {
	conn *ethclient.Client
	token *Token
}

// NewApi creates api handlers.
func NewApi() (*Api, error) {
	conn, err := ethclient.Dial("https://ropsten.infura.io/v3/a8d0d3bb333247d8b140a84938f9d414")
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the Ethereum client: %v", err)
	}

	token, err := NewToken(common.HexToAddress("0x4f7f1380239450AAD5af611DB3c3c1bb51049c29"), conn)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate a Token contract: %v", err)
	}

	return &Api{
		conn: conn,
		token: token,
	}, nil
}

// Groups
// @Summary fetches all groups ids.
// @Description get all groups
// @ID get-string-by-int
// @Produce  json
// @Success 200 {array} int
// @Failure 500 {object} httpError
// @Router /groups/ [get]
func (a *Api) Groups(c *fiber.Ctx) error {
	ids, err := a.token.GetGroupIds(&bind.CallOpts{})
	if err != nil {
		return fmt.Errorf("failed to retrieve group ids: %v", err)
	}

	return c.JSON(ids)
}

// GroupsById
// @Summary fetches the group information by id.
// @Description get group by ID
// @Produce  json
// @Param id path int true "Group ID"
// @Success 200 {object} nil
// @Failure 500 {object} httpError
// @Router /groups/{id} [get]
func (a *Api) GroupsById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return fmt.Errorf("id is not valid: %v", err)
	}

	group, err := a.token.GetGroup(&bind.CallOpts{}, big.NewInt(int64(id)))
	if err != nil {
		return fmt.Errorf("failed to retrieve group ids: %v", err)
	}

	return c.JSON(group)
}

// Indexes
// @Summary fetches the index information by id.
// @Description get index by ID
// @Produce  json
// @Param id path int true "Index ID"
// @Success 200 {object} nil
// @Failure 500 {object} httpError
// @Router /indexes/{id} [get]
func (a *Api) Indexes(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return fmt.Errorf("id is not valid: %v", err)
	}

	group, err := a.token.GetIndex(&bind.CallOpts{}, big.NewInt(int64(id)))
	if err != nil {
		return fmt.Errorf("failed to retrieve group ids: %v", err)
	}

	return c.JSON(group)
}

// BlockInfo
// @Summary get block header from the ropsten network by: latest, number, hash.
// @Produce  json
// @Param id path string true "Block info id"
// @Success 200 {object} nil
// @Failure 500 {object} httpError
// @Router /blocks/{id} [get]
func (a *Api) BlockInfo(c *fiber.Ctx) error {
	id := c.Params("id", "latest")

	// Get the latest block number info.
	if id == "latest"{
		n, err := a.conn.BlockNumber(c.Context())
		if err != nil {
			return fmt.Errorf("can't receive latest block number: %v", err)
		}
		blockInfo, err := a.conn.BlockByNumber(c.Context(), big.NewInt(int64(n)))
		if err != nil {
			return fmt.Errorf("can't receive latest block number: %v", err)
		}

		return c.JSON(blockInfo.Header())
	}

	// Get block info by the block number.
	if v, err := strconv.Atoi(id); err == nil {
		blockInfo, err := a.conn.BlockByNumber(c.Context(), big.NewInt(int64(v)))
		if err != nil {
			return fmt.Errorf("can't receive latest block number: %v", err)
		}

		return c.JSON(blockInfo.Header())
	}

	// Get block info by the hash number.
	blockInfo, err := a.conn.BlockByHash(c.Context(), common.HexToHash(id))
	if err != nil {
		return fmt.Errorf("can't receive latest block number: %v", err)
	}

	return c.JSON(blockInfo.Header())
}

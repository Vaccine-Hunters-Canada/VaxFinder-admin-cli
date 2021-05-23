package get

import (
	"context"
	"errors"
	"strconv"

	"vf-admin/internal/api"
	"vf-admin/internal/cmdrun/base"

	"github.com/spf13/cobra"
)

type requirementGetter struct{}

func (g requirementGetter) GetHeaders() []string {
	return []string{"id", "name", "description", "created at"}
}

func (g requirementGetter) GetResourceName() (string, string, string) {
	return "get", "got", "requirement"
}

func (g requirementGetter) GetArray(cmd *cobra.Command, args []string, client *api.ClientWithResponses) ([][]string, error) {
	id, aErr := strconv.Atoi(args[0])
	if aErr != nil {
		return nil, errors.New("expecting id as integer")
	}

	res, rErr := client.RetrieveRequirementByIdApiV1RequirementsRequirementIdGetWithResponse(context.Background(), id)
	if rErr != nil {
		return nil, rErr
	}

	if res.StatusCode() != 200 {
		return nil, errors.New(res.Status() + ": " + string(res.Body))
	}

	if res.JSON200 != nil {
		return [][]string{
			{
				strconv.Itoa(res.JSON200.Id), res.JSON200.Name, res.JSON200.Description, res.JSON200.CreatedAt.String(),
			},
		}, nil
	}
	return nil, nil
}

// CmdRunE is what's executed when running `vf-admin requirement get <id>`
func CmdRunE(cmd *cobra.Command, args []string) error {
	var getter requirementGetter

	base.Cmd(cmd, args, getter)
	return nil
}

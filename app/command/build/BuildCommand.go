package build

import (
	"deployRunner/command"
	"deployRunner/config"
	"errors"
	"fmt"
	"strings"
)

const (
	GetCrumbCommand   string = "curl -u %s:%s -s \"https://ci.platformtests.net/crumbIssuer/api/xml\" | xmllint --format --xpath \"concat(//crumbRequestField,':',//crumb)\" -"
	TriggerJobCommand string = "curl -I -X POST https://%s:%s@ci.platformtests.net/job/ecommpay/job/pp/job/concept-%s/view/tags/job/%s/build -H \"%s\""
)

type Command struct {
	params *command.ApplicationParams
	sdlc   *config.Sdlc
}

func New(application string, tag string, sdlc *config.Sdlc) Command {
	return Command{
		params: &command.ApplicationParams{Application: application, Tag: tag},
		sdlc:   sdlc,
	}
}

func (c Command) Run() error {
	crumb, err := command.Execute(fmt.Sprintf(GetCrumbCommand, c.sdlc.User, c.sdlc.Password), "")
	if err != nil {
		return err
	}

	response, err := command.Execute(fmt.Sprintf(TriggerJobCommand, c.sdlc.User, c.sdlc.Token, c.params.Application, c.params.Tag, crumb), "")
	if err != nil {
		return err
	}

	if strings.Contains(response, "404") {
		return errors.New("job does not exist")
	}

	return nil
}

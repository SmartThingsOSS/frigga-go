package frigga

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldDissectNameWithDot(t *testing.T) {
	names, err := Parse("chukwa.collector_1-v889")

	assert.Nil(t, err)
	assert.Equal(t, "chukwa.collector_1-v889", names.Group)
	assert.Equal(t, "chukwa.collector_1", names.Cluster)
	assert.Equal(t, "chukwa.collector_1", names.App)
	assert.Equal(t, "", names.Stack)
	assert.Equal(t, "", names.Detail)
	assert.Equal(t, "v889", names.Push)
	assert.Equal(t, "889", names.Sequence)
}

func TestShouldReturnErrorForInvalid(t *testing.T) {
	names, err := Parse("nccp-moviecontrol%27")

	assert.Nil(t, names)
	assert.Error(t, err)
}

func TestShouldDissectGroupNames(t *testing.T) {
	names, err := Parse("")

	assert.Error(t, err)
	assert.Nil(t, names)

	names, err = Parse("actiondrainer")

	assert.Nil(t, err)
	assert.Equal(t, "actiondrainer", names.Group)
	assert.Equal(t, "actiondrainer", names.Cluster)
	assert.Equal(t, "actiondrainer", names.App)
	assert.Equal(t, "", names.Stack)
	assert.Equal(t, "", names.Detail)
	assert.Equal(t, "", names.Push)
	assert.Equal(t, "", names.Sequence)

	names, err = Parse("actiondrainer-v003")

	assert.Nil(t, err)
	assert.Equal(t, "actiondrainer-v003", names.Group)
	assert.Equal(t, "actiondrainer", names.Cluster)
	assert.Equal(t, "actiondrainer", names.App)
	assert.Equal(t, "", names.Stack)
	assert.Equal(t, "", names.Detail)
	assert.Equal(t, "v003", names.Push)
	assert.Equal(t, "003", names.Sequence)

	names, err = Parse("actiondrainer--v003")

	assert.Nil(t, err)
	assert.Equal(t, "actiondrainer--v003", names.Group)
	assert.Equal(t, "actiondrainer-", names.Cluster)
	assert.Equal(t, "actiondrainer", names.App)
	assert.Equal(t, "", names.Stack)
	assert.Equal(t, "", names.Detail)
	assert.Equal(t, "v003", names.Push)
	assert.Equal(t, "003", names.Sequence)

	names, err = Parse("actiondrainer---v003")

	assert.Nil(t, err)
	assert.Equal(t, "actiondrainer---v003", names.Group)
	assert.Equal(t, "actiondrainer--", names.Cluster)
	assert.Equal(t, "actiondrainer", names.App)
	assert.Equal(t, "", names.Stack)
	assert.Equal(t, "", names.Detail)
	assert.Equal(t, "v003", names.Push)
	assert.Equal(t, "003", names.Sequence)

	names, err = Parse("api-test-A")

	assert.Nil(t, err)
	assert.Equal(t, "api-test-A", names.Group)
	assert.Equal(t, "api-test-A", names.Cluster)
	assert.Equal(t, "api", names.App)
	assert.Equal(t, "test", names.Stack)
	assert.Equal(t, "A", names.Detail)
	assert.Equal(t, "", names.Push)
	assert.Equal(t, "", names.Sequence)

	names, err = Parse("api-test-A-v406")

	assert.Nil(t, err)
	assert.Equal(t, "api-test-A-v406", names.Group)
	assert.Equal(t, "api-test-A", names.Cluster)
	assert.Equal(t, "api", names.App)
	assert.Equal(t, "test", names.Stack)
	assert.Equal(t, "A", names.Detail)
	assert.Equal(t, "v406", names.Push)
	assert.Equal(t, "406", names.Sequence)

	names, err = Parse("api-test101")

	assert.Nil(t, err)
	assert.Equal(t, "api-test101", names.Group)
	assert.Equal(t, "api-test101", names.Cluster)
	assert.Equal(t, "api", names.App)
	assert.Equal(t, "test101", names.Stack)
	assert.Equal(t, "", names.Detail)
	assert.Equal(t, "", names.Push)
	assert.Equal(t, "", names.Sequence)

	names, err = Parse("chukwacollector_1-v889")

	assert.Nil(t, err)
	assert.Equal(t, "chukwacollector_1-v889", names.Group)
	assert.Equal(t, "chukwacollector_1", names.Cluster)
	assert.Equal(t, "chukwacollector_1", names.App)
	assert.Equal(t, "", names.Stack)
	assert.Equal(t, "", names.Detail)
	assert.Equal(t, "v889", names.Push)
	assert.Equal(t, "889", names.Sequence)

	names, err = Parse("discovery-dev")

	assert.Nil(t, err)
	assert.Equal(t, "discovery-dev", names.Group)
	assert.Equal(t, "discovery-dev", names.Cluster)
	assert.Equal(t, "discovery", names.App)
	assert.Equal(t, "dev", names.Stack)
	assert.Equal(t, "", names.Detail)
	assert.Equal(t, "", names.Push)
	assert.Equal(t, "", names.Sequence)

	names, err = Parse("discovery-us-east-1d")

	assert.Nil(t, err)
	assert.Equal(t, "discovery-us-east-1d", names.Group)
	assert.Equal(t, "discovery-us-east-1d", names.Cluster)
	assert.Equal(t, "discovery", names.App)
	assert.Equal(t, "us", names.Stack)
	assert.Equal(t, "east-1d", names.Detail)
	assert.Equal(t, "", names.Push)
	assert.Equal(t, "", names.Sequence)

	names, err = Parse("evcache-us-east-1d-0")

	assert.Nil(t, err)
	assert.Equal(t, "evcache-us-east-1d-0", names.Group)
	assert.Equal(t, "evcache-us-east-1d-0", names.Cluster)
	assert.Equal(t, "evcache", names.App)
	assert.Equal(t, "us", names.Stack)
	assert.Equal(t, "east-1d-0", names.Detail)
	assert.Equal(t, "", names.Push)
	assert.Equal(t, "", names.Sequence)

	names, err = Parse("evcache-us-east-1d-0-v223")

	assert.Nil(t, err)
	assert.Equal(t, "evcache-us-east-1d-0-v223", names.Group)
	assert.Equal(t, "evcache-us-east-1d-0", names.Cluster)
	assert.Equal(t, "evcache", names.App)
	assert.Equal(t, "us", names.Stack)
	assert.Equal(t, "east-1d-0", names.Detail)
	assert.Equal(t, "v223", names.Push)
	assert.Equal(t, "223", names.Sequence)

	names, err = Parse("videometadata-navigator-integration-240-CAN")

	assert.Nil(t, err)
	assert.Equal(t, "videometadata-navigator-integration-240-CAN", names.Group)
	assert.Equal(t, "videometadata-navigator-integration-240-CAN", names.Cluster)
	assert.Equal(t, "videometadata", names.App)
	assert.Equal(t, "navigator", names.Stack)
	assert.Equal(t, "integration-240-CAN", names.Detail)
	assert.Equal(t, "", names.Push)
	assert.Equal(t, "", names.Sequence)
}

func TestDissectGroupNamesWithLabeledVariables(t *testing.T) {
	names, err := Parse("actiondrainer")

	assert.Nil(t, err)
	assert.Equal(t, "actiondrainer", names.Group)
	assert.Equal(t, "actiondrainer", names.Cluster)
	assert.Equal(t, "actiondrainer", names.App)
	assert.Equal(t, "", names.Stack)
	assert.Equal(t, "", names.Detail)
	assert.Equal(t, "", names.Push)
	assert.Equal(t, "", names.Sequence)
	assert.Equal(t, "", names.Countries)
	assert.Equal(t, "", names.DevPhase)
	assert.Equal(t, "", names.Hardware)
	assert.Equal(t, "", names.Partners)
	assert.Equal(t, "", names.Revision)
	assert.Equal(t, "", names.UsedBy)
	assert.Equal(t, "", names.RedBlackSwap)
	assert.Equal(t, "", names.Zone)

	names, err = Parse("cass-nccpintegration-random-junk-c0northamerica-d0prod-h0gamesystems-p0vizio-r027-u0nccp-w0A-z0useast1a-v003")

	assert.Nil(t, err)
	assert.Equal(t, "cass-nccpintegration-random-junk-c0northamerica-d0prod-h0gamesystems-p0vizio-r027-u0nccp-w0A-z0useast1a-v003", names.Group)
	assert.Equal(t, "cass-nccpintegration-random-junk-c0northamerica-d0prod-h0gamesystems-p0vizio-r027-u0nccp-w0A-z0useast1a", names.Cluster)
	assert.Equal(t, "cass", names.App)
	assert.Equal(t, "nccpintegration", names.Stack)
	assert.Equal(t, "random-junk", names.Detail)
	assert.Equal(t, "v003", names.Push)
	assert.Equal(t, "003", names.Sequence)
	assert.Equal(t, "northamerica", names.Countries)
	assert.Equal(t, "prod", names.DevPhase)
	assert.Equal(t, "gamesystems", names.Hardware)
	assert.Equal(t, "vizio", names.Partners)
	assert.Equal(t, "27", names.Revision)
	assert.Equal(t, "nccp", names.UsedBy)
	assert.Equal(t, "A", names.RedBlackSwap)
	assert.Equal(t, "useast1a", names.Zone)

	names, err = Parse("cass-nccpintegration-c0northamerica-d0prod")

	assert.Nil(t, err)
	assert.Equal(t, "cass-nccpintegration-c0northamerica-d0prod", names.Group)
	assert.Equal(t, "cass-nccpintegration-c0northamerica-d0prod", names.Cluster)
	assert.Equal(t, "cass", names.App)
	assert.Equal(t, "nccpintegration", names.Stack)
	assert.Equal(t, "", names.Detail)
	assert.Equal(t, "", names.Push)
	assert.Equal(t, "", names.Sequence)
	assert.Equal(t, "northamerica", names.Countries)
	assert.Equal(t, "prod", names.DevPhase)
	assert.Equal(t, "", names.Hardware)
	assert.Equal(t, "", names.Partners)
	assert.Equal(t, "", names.Revision)
	assert.Equal(t, "", names.UsedBy)
	assert.Equal(t, "", names.RedBlackSwap)
	assert.Equal(t, "", names.Zone)

	names, err = Parse("cass-c0northamerica-d0prod")

	assert.Nil(t, err)
	assert.Equal(t, "cass-c0northamerica-d0prod", names.Group)
	assert.Equal(t, "cass-c0northamerica-d0prod", names.Cluster)
	assert.Equal(t, "cass", names.App)
	assert.Equal(t, "", names.Stack)
	assert.Equal(t, "", names.Detail)
	assert.Equal(t, "", names.Push)
	assert.Equal(t, "", names.Sequence)
	assert.Equal(t, "northamerica", names.Countries)
	assert.Equal(t, "prod", names.DevPhase)
	assert.Equal(t, "", names.Hardware)
	assert.Equal(t, "", names.Partners)
	assert.Equal(t, "", names.Revision)
	assert.Equal(t, "", names.UsedBy)
	assert.Equal(t, "", names.RedBlackSwap)
	assert.Equal(t, "", names.Zone)

	names, err = Parse("cass-c0northamerica-d0prod-v102")

	assert.Nil(t, err)
	assert.Equal(t, "cass-c0northamerica-d0prod-v102", names.Group)
	assert.Equal(t, "cass-c0northamerica-d0prod", names.Cluster)
	assert.Equal(t, "cass", names.App)
	assert.Equal(t, "", names.Stack)
	assert.Equal(t, "", names.Detail)
	assert.Equal(t, "v102", names.Push)
	assert.Equal(t, "102", names.Sequence)
	assert.Equal(t, "northamerica", names.Countries)
	assert.Equal(t, "prod", names.DevPhase)
	assert.Equal(t, "", names.Hardware)
	assert.Equal(t, "", names.Partners)
	assert.Equal(t, "", names.Revision)
	assert.Equal(t, "", names.UsedBy)
	assert.Equal(t, "", names.RedBlackSwap)
	assert.Equal(t, "", names.Zone)

	names, err = Parse("cass-v102")

	assert.Nil(t, err)
	assert.Equal(t, "cass-v102", names.Group)
	assert.Equal(t, "cass", names.Cluster)
	assert.Equal(t, "cass", names.App)
	assert.Equal(t, "", names.Stack)
	assert.Equal(t, "", names.Detail)
	assert.Equal(t, "v102", names.Push)
	assert.Equal(t, "102", names.Sequence)
	assert.Equal(t, "", names.Countries)
	assert.Equal(t, "", names.DevPhase)
	assert.Equal(t, "", names.Hardware)
	assert.Equal(t, "", names.Partners)
	assert.Equal(t, "", names.Revision)
	assert.Equal(t, "", names.UsedBy)
	assert.Equal(t, "", names.RedBlackSwap)
	assert.Equal(t, "", names.Zone)
}

func TestExtractLabeledVariable(t *testing.T) {
	assert.Equal(t, "sony", extractLabeledVariable("-p0sony", "p"))
	assert.Equal(t, "northamerica", extractLabeledVariable("-c0northamerica-d0prod-h0gamesystems-p0vizio-r027-u0nccp-x0A-z0useast1a", "c"))
	assert.Equal(t, "prod", extractLabeledVariable("-c0northamerica-d0prod-h0gamesystems-p0vizio-r027-u0nccp-x0A-z0useast1a", "d"))
	assert.Equal(t, "gamesystems", extractLabeledVariable("-c0northamerica-d0prod-h0gamesystems-p0vizio-r027-u0nccp-x0A-z0useast1a", "h"))
	assert.Equal(t, "vizio", extractLabeledVariable("-c0northamerica-d0prod-h0gamesystems-p0vizio-r027-u0nccp-x0A-z0useast1a", "p"))
	assert.Equal(t, "27", extractLabeledVariable("-c0northamerica-d0prod-h0gamesystems-p0vizio-r027-u0nccp-x0A-z0useast1a", "r"))
	assert.Equal(t, "nccp", extractLabeledVariable("-c0northamerica-d0prod-h0gamesystems-p0vizio-r027-u0nccp-x0A-z0useast1a", "u"))
	assert.Equal(t, "A", extractLabeledVariable("-c0northamerica-d0prod-h0gamesystems-p0vizio-r027-u0nccp-x0A-z0useast1a", "x"))
	assert.Equal(t, "useast1a", extractLabeledVariable("-c0northamerica-d0prod-h0gamesystems-p0vizio-r027-u0nccp-x0A-z0useast1a", "z"))
	assert.Equal(t, "", extractLabeledVariable("-c0northamerica-d0prod-h0gamesystems-p0vizio-r027-u0nccp-x0A-z0useast1a", "a"))
	assert.Equal(t, "", extractLabeledVariable("", "a"))
	assert.Equal(t, "", extractLabeledVariable("", ""))
	assert.Equal(t, "", extractLabeledVariable("-c0northamerica-d0prod-h0gamesystems-p0vizio-r027-u0nccp-x0A-z0useast1a", ""))
	assert.Equal(t, "", extractLabeledVariable("-p0sony", ""))
}

func TestCaret(t *testing.T) {
    names, err := Parse("caret-prod-^1.0.0-v002")
    assert.Nil(t, err)
    assert.Equal(t, "caret-prod-^1.0.0-v002", names.Group)
    assert.Equal(t, "caret-prod-^1.0.0", names.Cluster)
    assert.Equal(t, "caret", names.App)
    assert.Equal(t, "prod", names.Stack)
    assert.Equal(t, "^1.0.0", names.Detail)
    assert.Equal(t, "v002", names.Push)
    assert.Equal(t, "002", names.Sequence)
}

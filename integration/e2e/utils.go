package e2e

import (
	"github.com/hyperledger/fabric/integration/nwo"
)

type ChaincodeQueryWithHex struct {
	ChannelID string
	Name      string
	Ctor      string
}

func (c ChaincodeQueryWithHex) SessionName() string {
	return "peer-chaincode-query"
}

func (c ChaincodeQueryWithHex) Args() []string {
	return []string{
		"chaincode", "query",
		"--channelID", c.ChannelID,
		"--name", c.Name,
		"--ctor", c.Ctor,
		"--hex",
	}
}

func Build() *nwo.Components {
	components := &nwo.Components{}
	components.Build()
	return components
}

func SimpleSoloNetwork() *nwo.Config {
	return &nwo.Config{
		Organizations: []*nwo.Organization{{
			Name:          "OrdererOrg",
			MSPID:         "OrdererMSP",
			Domain:        "example.com",
			EnableNodeOUs: false,
			Users:         0,
			CA:            &nwo.CA{Hostname: "ca"},
		}, {
			Name:          "Org1",
			MSPID:         "Org1MSP",
			Domain:        "org1.example.com",
			EnableNodeOUs: true,
			Users:         2,
			CA:            &nwo.CA{Hostname: "ca"},
		}},
		Consortiums: []*nwo.Consortium{{
			Name: "SampleConsortium",
			Organizations: []string{
				"Org1",
			},
		}},
		Consensus: &nwo.Consensus{
			Type: "solo",
		},
		SystemChannel: &nwo.SystemChannel{
			Name:    "systemchannel",
			Profile: "OneOrgOrdererGenesis",
		},
		Orderers: []*nwo.Orderer{
			{Name: "orderer", Organization: "OrdererOrg"},
		},
		Channels: []*nwo.Channel{
			{Name: "testchannel", Profile: "OneOrgChannel"},
		},
		Peers: []*nwo.Peer{{
			Name:         "peer0",
			Organization: "Org1",
			Channels: []*nwo.PeerChannel{
				{Name: "testchannel", Anchor: true},
			},
		}},
		Profiles: []*nwo.Profile{{
			Name:     "OneOrgOrdererGenesis",
			Orderers: []string{"orderer"},
		}, {
			Name:          "OneOrgChannel",
			Consortium:    "SampleConsortium",
			Organizations: []string{"Org1"},
		}},
	}
}

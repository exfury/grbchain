package transfer_test

import (
	"math"
	"testing"

	capabilitytypes "github.com/exfury/grbchain/libs/cosmos-sdk/x/capability/types"
	"github.com/exfury/grbchain/libs/ibc-go/modules/apps/transfer/types"
	channeltypes "github.com/exfury/grbchain/libs/ibc-go/modules/core/04-channel/types"
	host "github.com/exfury/grbchain/libs/ibc-go/modules/core/24-host"
	ibctesting "github.com/exfury/grbchain/libs/ibc-go/testing"
	"github.com/stretchr/testify/suite"
)

func (suite *TransferTestSuite) TestOnChanOpenInit() {
	var (
		channel *channeltypes.Channel
		path    *ibctesting.Path
		chanCap *capabilitytypes.Capability
	)

	testCases := []struct {
		name     string
		malleate func()
		expPass  bool
	}{

		{
			"success", func() {}, true,
		},
		{
			"max channels reached", func() {
				path.EndpointA.ChannelID = channeltypes.FormatChannelIdentifier(math.MaxUint32 + 1)
			}, false,
		},
		{
			"invalid order - ORDERED", func() {
				channel.Ordering = channeltypes.ORDERED
			}, false,
		},
		{
			"invalid port ID", func() {
				path.EndpointA.ChannelConfig.PortID = ibctesting.MockPort
			}, false,
		},
		{
			"invalid version", func() {
				channel.Version = "version"
			}, false,
		},
		{
			"capability already claimed", func() {
				err := suite.chainA.GetSimApp().ScopedTransferKeeper.ClaimCapability(suite.chainA.GetContext(), chanCap, host.ChannelCapabilityPath(path.EndpointA.ChannelConfig.PortID, path.EndpointA.ChannelID))
				suite.Require().NoError(err)
			}, false,
		},
	}

	for _, tc := range testCases {
		tc := tc

		suite.Run(tc.name, func() {
			suite.SetupTest() // reset
			path = NewTransferPath(suite.chainA, suite.chainB)
			suite.coordinator.SetupConnections(path)
			path.EndpointA.ChannelID = ibctesting.FirstChannelID

			counterparty := channeltypes.NewCounterparty(path.EndpointA.ChannelConfig.PortID, path.EndpointA.ChannelID)
			channel = &channeltypes.Channel{
				State:          channeltypes.INIT,
				Ordering:       channeltypes.UNORDERED,
				Counterparty:   counterparty,
				ConnectionHops: []string{path.EndpointA.ConnectionID},
				Version:        types.Version,
			}

			module, _, err := suite.chainA.App().GetIBCKeeper().PortKeeper.LookupModuleByPort(suite.chainA.GetContext(), ibctesting.TransferPort)
			suite.Require().NoError(err)

			chanCap, err = suite.chainA.App().GetScopedIBCKeeper().NewCapability(suite.chainA.GetContext(), host.ChannelCapabilityPath(ibctesting.TransferPort, path.EndpointA.ChannelID))
			suite.Require().NoError(err)

			cbs, ok := suite.chainA.App().GetIBCKeeper().Router.GetRoute(module)
			suite.Require().True(ok)

			tc.malleate() // explicitly change fields in channel and testChannel

			_, err = cbs.OnChanOpenInit(suite.chainA.GetContext(), channel.Ordering, channel.GetConnectionHops(),
				path.EndpointA.ChannelConfig.PortID, path.EndpointA.ChannelID, chanCap, channel.Counterparty, channel.GetVersion(),
			)

			if tc.expPass {
				suite.Require().NoError(err)
			} else {
				suite.Require().Error(err)
			}

		})
	}
}

func (suite *TransferTestSuite) TestOnChanOpenTry() {
	var (
		channel             *channeltypes.Channel
		chanCap             *capabilitytypes.Capability
		path                *ibctesting.Path
		counterpartyVersion string
	)

	testCases := []struct {
		name     string
		malleate func()
		expPass  bool
	}{
		//
		{
			"success", func() {}, true,
		},
		{
			"max channels reached", func() {
				path.EndpointA.ChannelID = channeltypes.FormatChannelIdentifier(math.MaxUint32 + 1)
			}, false,
		},
		{
			"capability already claimed in INIT should pass", func() {
				err := suite.chainA.GetSimApp().ScopedTransferKeeper.ClaimCapability(suite.chainA.GetContext(), chanCap, host.ChannelCapabilityPath(path.EndpointA.ChannelConfig.PortID, path.EndpointA.ChannelID))
				suite.Require().NoError(err)
			}, true,
		},
		{
			"invalid order - ORDERED", func() {
				channel.Ordering = channeltypes.ORDERED
			}, false,
		},
		{
			"invalid port ID", func() {
				path.EndpointA.ChannelConfig.PortID = ibctesting.MockPort
			}, false,
		},
		{
			"invalid counterparty version", func() {
				counterpartyVersion = "version"
			}, false,
		},
	}

	for _, tc := range testCases {
		tc := tc

		suite.Run(tc.name, func() {
			suite.SetupTest() // reset

			path = NewTransferPath(suite.chainA, suite.chainB)
			suite.coordinator.SetupConnections(path)
			path.EndpointA.ChannelID = ibctesting.FirstChannelID

			counterparty := channeltypes.NewCounterparty(path.EndpointA.ChannelConfig.PortID, path.EndpointA.ChannelID)
			channel = &channeltypes.Channel{
				State:          channeltypes.TRYOPEN,
				Ordering:       channeltypes.UNORDERED,
				Counterparty:   counterparty,
				ConnectionHops: []string{path.EndpointA.ConnectionID},
				Version:        types.Version,
			}
			counterpartyVersion = types.Version

			module, _, err := suite.chainA.App().GetIBCKeeper().PortKeeper.LookupModuleByPort(suite.chainA.GetContext(), ibctesting.TransferPort)
			suite.Require().NoError(err)

			chanCap, err = suite.chainA.App().GetScopedIBCKeeper().NewCapability(suite.chainA.GetContext(), host.ChannelCapabilityPath(ibctesting.TransferPort, path.EndpointA.ChannelID))
			suite.Require().NoError(err)

			cbs, ok := suite.chainA.App().GetIBCKeeper().Router.GetRoute(module)
			suite.Require().True(ok)

			tc.malleate() // explicitly change fields in channel and testChannel

			version, err := cbs.OnChanOpenTry(suite.chainA.GetContext(), channel.Ordering, channel.GetConnectionHops(),
				path.EndpointA.ChannelConfig.PortID, path.EndpointA.ChannelID, chanCap, channel.Counterparty, channel.GetVersion(), counterpartyVersion,
			)

			if tc.expPass {
				suite.Require().NoError(err)
				suite.Require().Equal(types.Version, version)
			} else {
				suite.Require().Error(err)
			}

		})
	}
}

func (suite *TransferTestSuite) TestOnChanOpenAck() {
	var counterpartyVersion string

	testCases := []struct {
		name     string
		malleate func()
		expPass  bool
	}{

		{
			"success", func() {}, true,
		},
		{
			"invalid counterparty version", func() {
				counterpartyVersion = "version"
			}, false,
		},
	}

	for _, tc := range testCases {
		tc := tc

		suite.Run(tc.name, func() {
			suite.SetupTest() // reset

			path := NewTransferPath(suite.chainA, suite.chainB)
			suite.coordinator.SetupConnections(path)
			path.EndpointA.ChannelID = ibctesting.FirstChannelID
			counterpartyVersion = types.Version

			module, _, err := suite.chainA.App().GetIBCKeeper().PortKeeper.LookupModuleByPort(suite.chainA.GetContext(), ibctesting.TransferPort)
			suite.Require().NoError(err)

			cbs, ok := suite.chainA.App().GetIBCKeeper().Router.GetRoute(module)
			suite.Require().True(ok)

			tc.malleate() // explicitly change fields in channel and testChannel

			err = cbs.OnChanOpenAck(suite.chainA.GetContext(), path.EndpointA.ChannelConfig.PortID, path.EndpointA.ChannelID, path.EndpointA.Counterparty.ChannelID, counterpartyVersion)

			if tc.expPass {
				suite.Require().NoError(err)
			} else {
				suite.Require().Error(err)
			}

		})
	}
}

func TestTransferTestSuite2(t *testing.T) {
	suite.Run(t, new(TransferTestSuite))
}
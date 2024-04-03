package keeper_test

import (
	testutil "github.com/jackalLabs/canine-chain/v3/testutil"
	"github.com/jackalLabs/canine-chain/v3/x/storage/types"
)

func (suite *KeeperTestSuite) TestReport() {
	suite.SetupSuite()
	params := suite.storageKeeper.GetParams(suite.ctx)

	addresses, err := testutil.CreateTestAddresses("cosmos", int(params.AttestFormSize)+2)
	suite.Require().NoError(err)

	badAddresses, err := testutil.CreateTestAddresses("cosmos", 10)

	cases := map[string]struct {
		owner   string
		merkle  []byte
		start   int64
		creator string
		expErr  bool
	}{
		"attestation form found": {
			merkle:  []byte("merkle"),
			owner:   "owner",
			start:   0,
			creator: addresses[0],
			expErr:  false,
		},
		"not requested provider": {
			merkle:  []byte("merkle"),
			owner:   "owner",
			start:   0,
			creator: badAddresses[9],
			expErr:  true,
		},
		"active deal not found": {
			merkle:  []byte("merkle_bad"),
			owner:   "owner",
			start:   0,
			creator: addresses[0],
			expErr:  true,
		},
	}

	for name, tc := range cases {
		suite.Run(name, func() {
			suite.SetupSuite()

			uf := types.UnifiedFile{
				Merkle:        []byte("merkle"),
				Owner:         "owner",
				Start:         0,
				Expires:       0,
				FileSize:      1000,
				ProofInterval: 100,
				ProofType:     0,
				Proofs:        make([]string, 0),
				MaxProofs:     3,
				Note:          "test",
			}
			suite.storageKeeper.SetFile(suite.ctx, uf)
			uf.AddProver(suite.ctx, suite.storageKeeper, addresses[4])

			attestations := make([]*types.Attestation, params.AttestFormSize)

			for i := 0; i < int(params.AttestFormSize); i++ {
				attestations[i] = &types.Attestation{
					Provider: addresses[i],
					Complete: false,
				}
			}

			attestForm := types.ReportForm{
				Attestations: attestations,
				Prover:       addresses[4],
				Merkle:       uf.Merkle,
				Owner:        uf.Owner,
				Start:        uf.Start,
			}

			noActiveDealAttestForm := types.ReportForm{
				Attestations: attestations,
				Prover:       addresses[4],
				Merkle:       []byte("no_merkle"),
				Owner:        uf.Owner,
				Start:        uf.Start,
			}

			suite.storageKeeper.SetReportForm(suite.ctx, attestForm)
			suite.storageKeeper.SetReportForm(suite.ctx, noActiveDealAttestForm)

			err = suite.storageKeeper.DoReport(suite.ctx, addresses[4], tc.merkle, tc.owner, tc.start, tc.creator)

			if tc.expErr {
				suite.Error(err)
			} else {
				suite.NoError(err)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestRequestReportForm() {
	suite.SetupSuite()
	params := suite.storageKeeper.GetParams(suite.ctx)

	addresses, err := testutil.CreateTestAddresses("cosmos", int(params.AttestFormSize)+10)
	suite.Require().NoError(err)

	cases := map[string]struct {
		owner   string
		merkle  []byte
		start   int64
		creator string
		expErr  bool
	}{
		"attestation already requested": {
			merkle:  []byte("merkle"),
			owner:   "owner",
			start:   0,
			creator: addresses[params.AttestFormSize],
			expErr:  true,
		},
		"not provider's cid": {
			merkle:  []byte("merkle"),
			owner:   "owner",
			start:   0,
			creator: "not provider's cid",
			expErr:  true,
		},
		"active deal not found": {
			merkle:  []byte("merkle_bad"),
			owner:   "owner",
			start:   0,
			creator: addresses[params.AttestFormSize],
			expErr:  true,
		},
	}

	for name, tc := range cases {
		suite.Run(name, func() {
			suite.SetupSuite()

			uf := types.UnifiedFile{
				Merkle:        []byte("merkle"),
				Owner:         "owner",
				Start:         0,
				Expires:       0,
				FileSize:      1000,
				ProofInterval: 100,
				ProofType:     0,
				Proofs:        make([]string, 0),
				MaxProofs:     3,
				Note:          "test",
			}
			suite.storageKeeper.SetFile(suite.ctx, uf)
			uf.AddProver(suite.ctx, suite.storageKeeper, addresses[10])

			for i := 0; i < int(params.AttestFormSize); i++ {
				provider := types.ActiveProviders{
					Address: addresses[i],
				}
				suite.storageKeeper.SetActiveProviders(suite.ctx, provider)
			}

			attestations := make([]*types.Attestation, params.AttestFormSize)

			for i := 0; i < int(params.AttestFormSize); i++ {
				attestations[i] = &types.Attestation{
					Provider: addresses[i],
					Complete: false,
				}
			}

			attestForm := types.ReportForm{
				Attestations: attestations,
				Prover:       addresses[10],
				Merkle:       []byte("no_merkle"),
				Owner:        uf.Owner,
				Start:        uf.Start,
			}

			suite.storageKeeper.SetReportForm(suite.ctx, attestForm)

			_, err = suite.storageKeeper.RequestReport(suite.ctx, addresses[10], tc.merkle, tc.owner, tc.start)

			if tc.expErr {
				suite.Error(err)
			} else {
				suite.NoError(err)
			}
		})
	}
}

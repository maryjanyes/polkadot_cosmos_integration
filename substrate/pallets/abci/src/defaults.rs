pub const DEFAULT_ABCI_URL: &str = "tcp://localhost:26658";

pub const DEFAULT_ABCI_APP_STATE: &str = r#"{
    "bank": {
        "params": {
            "default_send_enabled": true
        },
        "balances": [],
        "supply": [],
        "denom_metadata": [
            {
                "description": "The native staking token of the Cosmos Hub.",
                "denom_units": [
                    {
                        "denom": "uatom",
                        "exponent": 0,
                        "aliases": [
                            "microatom"
                        ]
                    },
                    {
                        "denom": "matom",
                        "exponent": 3,
                        "aliases": [
                            "milliatom"
                        ]
                    },
                    {
                        "denom": "atom",
                        "exponent": 6
                    }
                ],
                "base": "uatom",
                "display": "atom"
            }
        ]
    },
    "genutil": {
        "gentxs": []
    },
    "capability": {
        "index": "1",
        "owners": []
    },
    "mint": {
        "minter": {
            "inflation": "0.130000000000000000",
            "annual_provisions": "0.000000000000000000"
        },
        "params": {
            "mint_denom": "stake",
            "inflation_rate_change": "0.130000000000000000",
            "inflation_max": "0.200000000000000000",
            "inflation_min": "0.070000000000000000",
            "goal_bonded": "0.670000000000000000",
            "blocks_per_year": "6311520"
        }
    },
    "ibc": {
        "client_genesis": {
            "clients": [],
            "clients_consensus": [],
            "create_localhost": true
        },
        "connection_genesis": {
            "connections": [],
            "client_connection_paths": []
        },
        "channel_genesis": {
            "channels": [],
            "acknowledgements": [],
            "commitments": [],
            "send_sequences": [],
            "recv_sequences": [],
            "ack_sequences": []
        }
    },
    "upgrade": {},
    "evidence": {
        "evidence": []
    },
    "auth": {
        "params": {
            "max_memo_characters": "256",
            "tx_sig_limit": "7",
            "tx_size_cost_per_byte": "10",
            "sig_verify_cost_ed25519": "590",
            "sig_verify_cost_secp256k1": "1000"
        },
        "accounts": []
    },
    "gov": {
        "starting_proposal_id": "1",
        "deposits": null,
        "votes": null,
        "proposals": null,
        "deposit_params": {
            "min_deposit": [
                {
                    "denom": "stake",
                    "amount": "10000000"
                }
            ],
            "max_deposit_period": "172800000000000"
        },
        "voting_params": {
            "voting_period": "172800000000000"
        },
        "tally_params": {
            "quorum": "0.334000000000000000",
            "threshold": "0.500000000000000000",
            "veto": "0.334000000000000000"
        }
    },
    "params": null,
    "transfer": {
        "port_id": "transfer"
    },
    "crisis": {
        "constant_fee": {
            "denom": "stake",
            "amount": "1000"
        }
    },
    "distribution": {
        "params": {
            "community_tax": "0.020000000000000000",
            "base_proposer_reward": "0.010000000000000000",
            "bonus_proposer_reward": "0.040000000000000000",
            "withdraw_addr_enabled": true
        },
        "fee_pool": {
            "community_pool": []
        },
        "delegator_withdraw_infos": [],
        "previous_proposer": "",
        "outstanding_rewards": [],
        "validator_accumulated_commissions": [],
        "validator_historical_rewards": [],
        "validator_current_rewards": [],
        "delegator_starting_infos": [],
        "validator_slash_events": []
    },
    "slashing": {
        "params": {
            "signed_blocks_window": "100",
            "min_signed_per_window": "0.500000000000000000",
            "downtime_jail_duration": "600000000000",
            "slash_fraction_double_sign": "0.050000000000000000",
            "slash_fraction_downtime": "0.010000000000000000"
        },
        "signing_infos": {},
        "missed_blocks": {}
    },
    "staking": {
        "params": {
            "unbonding_time": "1814400000000000",
            "max_validators": 100,
            "max_entries": 7,
            "historical_entries": 100,
            "bond_denom": "stake"
        },
        "last_total_power": "0",
        "last_validator_powers": null,
        "validators": null,
        "delegations": null,
        "unbonding_delegations": null,
        "redelegations": null,
        "exported": false
    }
}"#;

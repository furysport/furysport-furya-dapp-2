/**
* This file was automatically generated by @cosmwasm/ts-codegen@0.16.3.
* DO NOT MODIFY IT BY HAND. Instead, modify the source JSONSchema file,
* and run the @cosmwasm/ts-codegen generate command to regenerate this file.
*/

export type Addr = string;
export interface ConfigResponse {
  denom: string;
  last_report_id: number;
  owner: Addr;
  [k: string]: unknown;
}
export interface Config {
  denom: string;
  last_report_id: number;
  owner: Addr;
  [k: string]: unknown;
}
export type Cw721HookMsg = {
  stake: {
    [k: string]: unknown;
  };
};
export type ExecuteMsg = {
  update_config: {
    owner: string;
    [k: string]: unknown;
  };
} | {
  add_daily_report: {
    report_id: number;
    rewards: RewardInfo[];
    [k: string]: unknown;
  };
} | {
  claim: {
    [k: string]: unknown;
  };
};
export type Uint128 = string;
export interface RewardInfo {
  addr: string;
  amount: Uint128;
  [k: string]: unknown;
}
export interface InstantiateMsg {
  denom: string;
  [k: string]: unknown;
}
export type QueryMsg = {
  config: {
    [k: string]: unknown;
  };
} | {
  total_claimable: {
    [k: string]: unknown;
  };
} | {
  user_claimable: {
    addr: string;
    [k: string]: unknown;
  };
} | {
  report_amount: {
    addr: string;
    report_id: number;
    [k: string]: unknown;
  };
};
export type ReportAmountResponse = string;
export type TotalClaimableResponse = string;
export type UserClaimableResponse = string;
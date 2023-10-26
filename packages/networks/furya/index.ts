import { furyaCurrencies } from "./currencies";
import { NetworkKind, CosmosNetworkInfo, NetworkFeature } from "../types";

const nameServiceContractAddress =
  "furya1wkwy0xh89ksdgj9hr347dyd2dw7zesmtrue6kfzyml4vdtz6e5ws7vvsy7";
const riotContractAddressGen1 =
  "furya1gflccmghzfscmxl95z43v36y0rle8v9x8kvt9na03yzywtw86ams5epky8";

export const furyaNetwork: CosmosNetworkInfo = {
  id: "furya",
  kind: NetworkKind.Cosmos,
  chainId: "furya-1",
  displayName: "Furya",
  icon: "icons/networks/furya.svg",
  features: [
    NetworkFeature.NFTMarketplace,
    NetworkFeature.Organizations,
    NetworkFeature.SocialFeed,
    NetworkFeature.UPP,
    NetworkFeature.NameService,
    NetworkFeature.BurnTokens,
    NetworkFeature.NFTLaunchpad,
    NetworkFeature.RiotP2E,
  ],
  overrides: "cosmos-registry:furya",
  walletUrlForStaking: "https://app.furya.xyz/staking",
  currencies: furyaCurrencies,
  txExplorer: "https://www.mintscan.io/furya/txs/$hash",
  accountExplorer: "https://www.mintscan.io/furya/account/$address",
  contractExplorer: "https://www.mintscan.io/furya/account/$address",
  idPrefix: "furya",
  testnet: false,
  backendEndpoint: "https://dapp-backend.mainnet.furya.xyz",
  addressPrefix: "furya",
  restEndpoint: "https://api.furya.xyz",
  rpcEndpoint: "https://rpc.furya.xyz",
  stakeCurrency: "ufury",
  gasPriceStep: {
    low: 0.0,
    average: 0.025,
    high: 0.04,
  },
  cosmosFeatures: [
    "stargate",
    "ibc-transfer",
    "cosmwasm",
    "no-legacy-stdTx",
    "ibc-go",
  ],
  nameServiceContractAddress,
  nameServiceDefaultImage:
    "ipfs://bafkreieqcwmjcb64r42ygs6a4dswz63djzgayjn3rhzjber3e42cknawlm",
  nameServiceTLD: ".furya",
  vaultContractAddress:
    "furya14hj2tavq8fpesdwxxcu44rty3hh90vhujrvcmstl4zr3txmfvw9shtt33f",
  riotContractAddressGen0:
    "furya1mf6ptkssddfmxvhdx0ech0k03ktp6kf9yk59renau2gvht3nq2gqwmv5rh",
  riotContractAddressGen1,
  riotSquadStakingContractAddressV1:
    "furya1zwkmdfsc6h96jm4yqdykaw5y7ndwsvqvgh5ryxp9nxs3hccml0xq0c3yy8",
  riotSquadStakingContractAddressV2:
    "furya1kvjk9m7dk0es35y6ah0k28llllvle3n7xgvh0gh568ta0paf8awsarcev4",
  distributorContractAddress:
    "furya1mnem0rhjaxcsghe0xw692xysra63xwwdee2wth9s0rkfwh4dxpaqcxjlkq",
  riotersFooterContractAddress:
    "furya1m6g3l3j5t988zwmp965hrhsxvm8jrkf2ulw4gmwj8hx8pd84tvcql90u95",
  secondaryDuringMintList: [
    nameServiceContractAddress,
    riotContractAddressGen1,
    "furya1lnx4r7styl209e9lfce8tdd7hyclq98upx25ax3t2qkmcl3jlgvs3z4p4y", // Glitch Candies
    "furya167xst2jy9n6u92t3n8hf762adtpe3cs6acsgn0w5n2xlz9hv3xgsn2z80j", // Diseases of the Brain AI
  ],
  excludeFromLaunchpadList: [riotContractAddressGen1],
  socialFeedContractAddress:
    "furya1lxf8agg0wd2m7n2ultl0yx337jw23puh0mlkkw5vhtnkkfettwfqzpvf5c",
  daoFactoryContractAddress:
    "furya16rxh5hgukhdq8rvm2j3t6v483dcqguwp4l825vlwz5pmfpw7s4rsngpvdy",
  daoCoreCodeId: 27,
  daoPreProposeSingleCodeId: 28,
  daoProposalSingleCodeId: 31,
  daoCw4GroupCodeId: 32,
  daoVotingCw4CodeId: 33,
};

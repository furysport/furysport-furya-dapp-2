import { useQuery } from "@tanstack/react-query";

import { FuryaMinter__factory } from "../../evm-contracts-clients/furya-bunker-minter/FuryaMinter__factory";
import { NetworkKind, getNetwork } from "../../networks";
import { getEthereumProvider } from "../../utils/ethereum";

export const useEthMinterIsPaused = (
  networkId: string | undefined,
  mintAddress: string,
  enabled?: boolean
) => {
  if (enabled === undefined) {
    enabled = true;
  }
  return useQuery(
    ["ethMinterIsPaused", networkId, mintAddress],
    async () => {
      const network = getNetwork(networkId);
      if (network?.kind !== NetworkKind.Ethereum) {
        return undefined;
      }

      const provider = await getEthereumProvider(network);
      if (!provider) {
        return undefined;
      }

      const minterClient = FuryaMinter__factory.connect(
        mintAddress,
        provider
      );
      return await minterClient.callStatic.paused();
    },
    { staleTime: Infinity, enabled: enabled && !!networkId }
  );
};

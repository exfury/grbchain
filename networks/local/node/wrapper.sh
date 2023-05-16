#!/usr/bin/env sh

##
## Input parameters
##
ID=${ID:-0}
LOG=${LOG:-grbchaind.log}

##
## Run binary with all parameters
##
export GRBCHAINDHOME="/grbchaind/node${ID}/okbchaind"

if [ -d "$(dirname "${GRBCHAINDHOME}"/"${LOG}")" ]; then
  grbchaind --chain-id okbchain-1 --home "${GRBCHAINDHOME}" "$@" | tee "${OKBCHAINDHOME}/${LOG}"
else
  grbchaind --chain-id okbchain-1 --home "${GRBCHAINDHOME}" "$@"
fi


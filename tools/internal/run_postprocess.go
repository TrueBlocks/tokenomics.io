package internal

import (
	"fmt"

	"github.com/spf13/cobra"
)

func RunPost(cmd *cobra.Command, args []string) error {
	fmt.Println("postprocess called")
	return nil
}

// STATEMENTS
// cat statements/$addr.csv | cut -d, -f1,2,3,4,5,6,9,25,26,30-33 | tee statements/balances/$addr.csv

// "assetAddress","assetSymbol","decimals","blockNumber","transactionIndex","timestamp","begBal","amountNet","endBal","reconciled","spotPrice","priceSource","reconciliationType"

// echo "count,assetAddr,assetSymbol" | tee statements/tx_counts/$addr.csv
// cat statements/balances/$addr.csv | \
//    grep -v assetAddr | \
//    cut -d, -f1,2 | \
//    sort | \
//    uniq -c | \
//    sort -n -r | \
//    sed 's/ //g' | \
//    sed 's/"/,/g' | \
//    cut -d, -f1,2,5
//    | tee -a statements/tx_counts/$addr.csv

// update the neighbor's images
// update the balance history charts

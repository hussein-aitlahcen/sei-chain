package antedecorators_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sei-protocol/sei-chain/app/antedecorators"
	"github.com/sei-protocol/sei-chain/utils"
	"github.com/stretchr/testify/require"
)

func TestTracedDecorator(t *testing.T) {
	output = ""
	anteDecorators := []sdk.AnteFullDecorator{
		sdk.DefaultWrappedAnteDecorator(FakeAnteDecoratorOne{}),
		sdk.DefaultWrappedAnteDecorator(FakeAnteDecoratorTwo{}),
		sdk.DefaultWrappedAnteDecorator(FakeAnteDecoratorThree{}),
	}
	tracedDecorators := utils.Map(anteDecorators, func(d sdk.AnteFullDecorator) sdk.AnteFullDecorator {
		return sdk.DefaultWrappedAnteDecorator(antedecorators.NewTracedAnteDecorator(d, nil))
	})
	chainedHandler, _ := sdk.ChainAnteDecorators(tracedDecorators...)
	chainedHandler(sdk.Context{}, FakeTx{}, false)
	require.Equal(t, "onetwothree", output)
}

package ssm

import (
	"github.com/stretchr/testify/assert"
	awsInit "main/aws"
	"testing"
)

func TestAwsGetParams(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		err := awsInit.InitAws("ap-northeast-2")
		assert.Nil(t, err)

		arr, err := AwsGetParams([]string{"unit_test", "unit_test_2"})
		assert.Nil(t, err)
		assert.Equal(t, "joker_ryan", arr[0])
		assert.Equal(t, "joker_ryan_2", arr[1])
	})
}

func TestAwsGetParam(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		err := awsInit.InitAws("ap-northeast-2")
		assert.Nil(t, err)

		value, err := AwsGetParam("unit_test")
		assert.Nil(t, err)
		assert.Equal(t, "joker_ryan", value)
	})
}

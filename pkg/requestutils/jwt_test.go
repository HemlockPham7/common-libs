package requestutils

import (
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
)

func TestGetUserIDFromRequest(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name string

		setupRequest func(ctx *gin.Context)

		expectedClaims string
		expectedErr    error
	}{
		{
			name: "claim not set in context",

			setupRequest: func(ctx *gin.Context) {

			},

			expectedClaims: "",
			expectedErr:    ErrNotExist,
		},
		{
			name: "claim is wrong type in context",

			setupRequest: func(ctx *gin.Context) {
				ctx.Set("claims", "string")
			},

			expectedClaims: "",
			expectedErr:    ErrInvalidToken,
		},
		{
			name: "claim not set in context",

			setupRequest: func(ctx *gin.Context) {
				ctx.Set("claims", jwt.MapClaims{
					"sub": "",
				})
			},

			expectedClaims: "",
			expectedErr:    ErrEmptyID,
		},
		{
			name: "valid claims in context",

			setupRequest: func(ctx *gin.Context) {
				ctx.Set("claims", jwt.MapClaims{
					"sub": "de305d54-75b4-431b-adb2-eb6b9e546099",
				})
			},

			expectedClaims: "de305d54-75b4-431b-adb2-eb6b9e546099",
			expectedErr:    nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			rec := httptest.NewRecorder()        // rec will store http status, response body, and header
			ctx, _ := gin.CreateTestContext(rec) // ctx.Writer = rec, rec is ResponseWriter, means every handler call c.JSON(...) would write the value into rec
			tc.setupRequest(ctx)

			uid, err := GetUserIDFromRequest(ctx)
			assert.Equal(t, tc.expectedErr, err)
			assert.Equal(t, tc.expectedClaims, uid)
		})
	}
}

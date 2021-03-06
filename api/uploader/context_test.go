package uploader_test

import (
	"testing"

	"github.com/cloudinary/cloudinary-go/api"
	"github.com/cloudinary/cloudinary-go/api/uploader"
	"github.com/cloudinary/cloudinary-go/internal/cldtest"
)

func TestUploader_Context(t *testing.T) {
	cldtest.UploadTestAsset(t, cldtest.PublicID)

	params := uploader.AddContextParams{
		PublicIDs: api.CldAPIArray{cldtest.PublicID},
		Context:   cldtest.CldContext,
	}

	resp, err := uploadAPI.AddContext(ctx, params)

	if err != nil {
		t.Error(err)
	}

	if resp == nil || len(resp.PublicIds) != 1 || resp.PublicIds[0] != cldtest.PublicID {
		t.Error(resp)
	}

	raParams := uploader.RemoveAllContextParams{
		PublicIDs: api.CldAPIArray{cldtest.PublicID},
	}

	raResp, err := uploadAPI.RemoveAllContext(ctx, raParams)

	if err != nil {
		t.Error(err)
	}

	if raResp == nil || len(raResp.PublicIds) != 1 || raResp.PublicIds[0] != cldtest.PublicID {
		t.Error(resp)
	}
}

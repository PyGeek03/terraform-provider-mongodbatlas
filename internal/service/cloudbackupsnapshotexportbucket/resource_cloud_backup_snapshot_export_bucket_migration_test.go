package cloudbackupsnapshotexportbucket_test

import (
	"testing"

	"github.com/mongodb/terraform-provider-mongodbatlas/internal/testutil/mig"
)

func TestMigBackupSnapshotExportBucket_basic(t *testing.T) {
	mig.SkipIfVersionBelow(t, "1.16.1")
	mig.CreateAndRunTest(t, basicTestCase(t))
}

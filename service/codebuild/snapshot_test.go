// Code generated by smithy-go-codegen DO NOT EDIT.

//go:build snapshot

package codebuild

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/smithy-go/middleware"
	"io"
	"io/fs"
	"os"
	"testing"
)

const ssprefix = "snapshot"

type snapshotOK struct{}

func (snapshotOK) Error() string { return "error: success" }

func createp(path string) (*os.File, error) {
	if err := os.Mkdir(ssprefix, 0700); err != nil && !errors.Is(err, fs.ErrExist) {
		return nil, err
	}
	return os.Create(path)
}

func sspath(op string) string {
	return fmt.Sprintf("%s/api_op_%s.go.snap", ssprefix, op)
}

func updateSnapshot(stack *middleware.Stack, operation string) error {
	f, err := createp(sspath(operation))
	if err != nil {
		return err
	}
	defer f.Close()
	if _, err := f.Write([]byte(stack.String())); err != nil {
		return err
	}
	return snapshotOK{}
}

func testSnapshot(stack *middleware.Stack, operation string) error {
	f, err := os.Open(sspath(operation))
	if errors.Is(err, fs.ErrNotExist) {
		return snapshotOK{}
	}
	if err != nil {
		return err
	}
	defer f.Close()
	expected, err := io.ReadAll(f)
	if err != nil {
		return err
	}
	if actual := stack.String(); actual != string(expected) {
		return fmt.Errorf("%s != %s", expected, actual)
	}
	return snapshotOK{}
}
func TestCheckSnapshot_BatchDeleteBuilds(t *testing.T) {
	svc := New(Options{})
	_, err := svc.BatchDeleteBuilds(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return testSnapshot(stack, "BatchDeleteBuilds")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestCheckSnapshot_BatchGetBuildBatches(t *testing.T) {
	svc := New(Options{})
	_, err := svc.BatchGetBuildBatches(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return testSnapshot(stack, "BatchGetBuildBatches")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestCheckSnapshot_BatchGetBuilds(t *testing.T) {
	svc := New(Options{})
	_, err := svc.BatchGetBuilds(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return testSnapshot(stack, "BatchGetBuilds")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestCheckSnapshot_BatchGetFleets(t *testing.T) {
	svc := New(Options{})
	_, err := svc.BatchGetFleets(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return testSnapshot(stack, "BatchGetFleets")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestCheckSnapshot_BatchGetProjects(t *testing.T) {
	svc := New(Options{})
	_, err := svc.BatchGetProjects(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return testSnapshot(stack, "BatchGetProjects")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestCheckSnapshot_BatchGetReportGroups(t *testing.T) {
	svc := New(Options{})
	_, err := svc.BatchGetReportGroups(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return testSnapshot(stack, "BatchGetReportGroups")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestCheckSnapshot_BatchGetReports(t *testing.T) {
	svc := New(Options{})
	_, err := svc.BatchGetReports(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return testSnapshot(stack, "BatchGetReports")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestCheckSnapshot_CreateFleet(t *testing.T) {
	svc := New(Options{})
	_, err := svc.CreateFleet(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return testSnapshot(stack, "CreateFleet")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestCheckSnapshot_CreateProject(t *testing.T) {
	svc := New(Options{})
	_, err := svc.CreateProject(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return testSnapshot(stack, "CreateProject")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestCheckSnapshot_CreateReportGroup(t *testing.T) {
	svc := New(Options{})
	_, err := svc.CreateReportGroup(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return testSnapshot(stack, "CreateReportGroup")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestCheckSnapshot_CreateWebhook(t *testing.T) {
	svc := New(Options{})
	_, err := svc.CreateWebhook(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return testSnapshot(stack, "CreateWebhook")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestCheckSnapshot_DeleteBuildBatch(t *testing.T) {
	svc := New(Options{})
	_, err := svc.DeleteBuildBatch(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return testSnapshot(stack, "DeleteBuildBatch")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestCheckSnapshot_DeleteFleet(t *testing.T) {
	svc := New(Options{})
	_, err := svc.DeleteFleet(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return testSnapshot(stack, "DeleteFleet")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestCheckSnapshot_DeleteProject(t *testing.T) {
	svc := New(Options{})
	_, err := svc.DeleteProject(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return testSnapshot(stack, "DeleteProject")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestCheckSnapshot_DeleteReport(t *testing.T) {
	svc := New(Options{})
	_, err := svc.DeleteReport(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return testSnapshot(stack, "DeleteReport")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestCheckSnapshot_DeleteReportGroup(t *testing.T) {
	svc := New(Options{})
	_, err := svc.DeleteReportGroup(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return testSnapshot(stack, "DeleteReportGroup")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestCheckSnapshot_DeleteResourcePolicy(t *testing.T) {
	svc := New(Options{})
	_, err := svc.DeleteResourcePolicy(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return testSnapshot(stack, "DeleteResourcePolicy")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestCheckSnapshot_DeleteSourceCredentials(t *testing.T) {
	svc := New(Options{})
	_, err := svc.DeleteSourceCredentials(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return testSnapshot(stack, "DeleteSourceCredentials")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestCheckSnapshot_DeleteWebhook(t *testing.T) {
	svc := New(Options{})
	_, err := svc.DeleteWebhook(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return testSnapshot(stack, "DeleteWebhook")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestCheckSnapshot_DescribeCodeCoverages(t *testing.T) {
	svc := New(Options{})
	_, err := svc.DescribeCodeCoverages(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return testSnapshot(stack, "DescribeCodeCoverages")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestCheckSnapshot_DescribeTestCases(t *testing.T) {
	svc := New(Options{})
	_, err := svc.DescribeTestCases(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return testSnapshot(stack, "DescribeTestCases")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestCheckSnapshot_GetReportGroupTrend(t *testing.T) {
	svc := New(Options{})
	_, err := svc.GetReportGroupTrend(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return testSnapshot(stack, "GetReportGroupTrend")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestCheckSnapshot_GetResourcePolicy(t *testing.T) {
	svc := New(Options{})
	_, err := svc.GetResourcePolicy(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return testSnapshot(stack, "GetResourcePolicy")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestCheckSnapshot_ImportSourceCredentials(t *testing.T) {
	svc := New(Options{})
	_, err := svc.ImportSourceCredentials(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return testSnapshot(stack, "ImportSourceCredentials")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestCheckSnapshot_InvalidateProjectCache(t *testing.T) {
	svc := New(Options{})
	_, err := svc.InvalidateProjectCache(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return testSnapshot(stack, "InvalidateProjectCache")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestCheckSnapshot_ListBuildBatches(t *testing.T) {
	svc := New(Options{})
	_, err := svc.ListBuildBatches(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return testSnapshot(stack, "ListBuildBatches")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestCheckSnapshot_ListBuildBatchesForProject(t *testing.T) {
	svc := New(Options{})
	_, err := svc.ListBuildBatchesForProject(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return testSnapshot(stack, "ListBuildBatchesForProject")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestCheckSnapshot_ListBuilds(t *testing.T) {
	svc := New(Options{})
	_, err := svc.ListBuilds(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return testSnapshot(stack, "ListBuilds")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestCheckSnapshot_ListBuildsForProject(t *testing.T) {
	svc := New(Options{})
	_, err := svc.ListBuildsForProject(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return testSnapshot(stack, "ListBuildsForProject")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestCheckSnapshot_ListCuratedEnvironmentImages(t *testing.T) {
	svc := New(Options{})
	_, err := svc.ListCuratedEnvironmentImages(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return testSnapshot(stack, "ListCuratedEnvironmentImages")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestCheckSnapshot_ListFleets(t *testing.T) {
	svc := New(Options{})
	_, err := svc.ListFleets(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return testSnapshot(stack, "ListFleets")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestCheckSnapshot_ListProjects(t *testing.T) {
	svc := New(Options{})
	_, err := svc.ListProjects(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return testSnapshot(stack, "ListProjects")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestCheckSnapshot_ListReportGroups(t *testing.T) {
	svc := New(Options{})
	_, err := svc.ListReportGroups(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return testSnapshot(stack, "ListReportGroups")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestCheckSnapshot_ListReports(t *testing.T) {
	svc := New(Options{})
	_, err := svc.ListReports(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return testSnapshot(stack, "ListReports")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestCheckSnapshot_ListReportsForReportGroup(t *testing.T) {
	svc := New(Options{})
	_, err := svc.ListReportsForReportGroup(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return testSnapshot(stack, "ListReportsForReportGroup")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestCheckSnapshot_ListSharedProjects(t *testing.T) {
	svc := New(Options{})
	_, err := svc.ListSharedProjects(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return testSnapshot(stack, "ListSharedProjects")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestCheckSnapshot_ListSharedReportGroups(t *testing.T) {
	svc := New(Options{})
	_, err := svc.ListSharedReportGroups(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return testSnapshot(stack, "ListSharedReportGroups")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestCheckSnapshot_ListSourceCredentials(t *testing.T) {
	svc := New(Options{})
	_, err := svc.ListSourceCredentials(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return testSnapshot(stack, "ListSourceCredentials")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestCheckSnapshot_PutResourcePolicy(t *testing.T) {
	svc := New(Options{})
	_, err := svc.PutResourcePolicy(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return testSnapshot(stack, "PutResourcePolicy")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestCheckSnapshot_RetryBuild(t *testing.T) {
	svc := New(Options{})
	_, err := svc.RetryBuild(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return testSnapshot(stack, "RetryBuild")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestCheckSnapshot_RetryBuildBatch(t *testing.T) {
	svc := New(Options{})
	_, err := svc.RetryBuildBatch(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return testSnapshot(stack, "RetryBuildBatch")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestCheckSnapshot_StartBuild(t *testing.T) {
	svc := New(Options{})
	_, err := svc.StartBuild(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return testSnapshot(stack, "StartBuild")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestCheckSnapshot_StartBuildBatch(t *testing.T) {
	svc := New(Options{})
	_, err := svc.StartBuildBatch(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return testSnapshot(stack, "StartBuildBatch")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestCheckSnapshot_StopBuild(t *testing.T) {
	svc := New(Options{})
	_, err := svc.StopBuild(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return testSnapshot(stack, "StopBuild")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestCheckSnapshot_StopBuildBatch(t *testing.T) {
	svc := New(Options{})
	_, err := svc.StopBuildBatch(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return testSnapshot(stack, "StopBuildBatch")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestCheckSnapshot_UpdateFleet(t *testing.T) {
	svc := New(Options{})
	_, err := svc.UpdateFleet(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return testSnapshot(stack, "UpdateFleet")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestCheckSnapshot_UpdateProject(t *testing.T) {
	svc := New(Options{})
	_, err := svc.UpdateProject(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return testSnapshot(stack, "UpdateProject")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestCheckSnapshot_UpdateProjectVisibility(t *testing.T) {
	svc := New(Options{})
	_, err := svc.UpdateProjectVisibility(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return testSnapshot(stack, "UpdateProjectVisibility")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestCheckSnapshot_UpdateReportGroup(t *testing.T) {
	svc := New(Options{})
	_, err := svc.UpdateReportGroup(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return testSnapshot(stack, "UpdateReportGroup")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestCheckSnapshot_UpdateWebhook(t *testing.T) {
	svc := New(Options{})
	_, err := svc.UpdateWebhook(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return testSnapshot(stack, "UpdateWebhook")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}
func TestUpdateSnapshot_BatchDeleteBuilds(t *testing.T) {
	svc := New(Options{})
	_, err := svc.BatchDeleteBuilds(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return updateSnapshot(stack, "BatchDeleteBuilds")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestUpdateSnapshot_BatchGetBuildBatches(t *testing.T) {
	svc := New(Options{})
	_, err := svc.BatchGetBuildBatches(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return updateSnapshot(stack, "BatchGetBuildBatches")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestUpdateSnapshot_BatchGetBuilds(t *testing.T) {
	svc := New(Options{})
	_, err := svc.BatchGetBuilds(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return updateSnapshot(stack, "BatchGetBuilds")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestUpdateSnapshot_BatchGetFleets(t *testing.T) {
	svc := New(Options{})
	_, err := svc.BatchGetFleets(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return updateSnapshot(stack, "BatchGetFleets")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestUpdateSnapshot_BatchGetProjects(t *testing.T) {
	svc := New(Options{})
	_, err := svc.BatchGetProjects(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return updateSnapshot(stack, "BatchGetProjects")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestUpdateSnapshot_BatchGetReportGroups(t *testing.T) {
	svc := New(Options{})
	_, err := svc.BatchGetReportGroups(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return updateSnapshot(stack, "BatchGetReportGroups")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestUpdateSnapshot_BatchGetReports(t *testing.T) {
	svc := New(Options{})
	_, err := svc.BatchGetReports(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return updateSnapshot(stack, "BatchGetReports")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestUpdateSnapshot_CreateFleet(t *testing.T) {
	svc := New(Options{})
	_, err := svc.CreateFleet(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return updateSnapshot(stack, "CreateFleet")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestUpdateSnapshot_CreateProject(t *testing.T) {
	svc := New(Options{})
	_, err := svc.CreateProject(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return updateSnapshot(stack, "CreateProject")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestUpdateSnapshot_CreateReportGroup(t *testing.T) {
	svc := New(Options{})
	_, err := svc.CreateReportGroup(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return updateSnapshot(stack, "CreateReportGroup")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestUpdateSnapshot_CreateWebhook(t *testing.T) {
	svc := New(Options{})
	_, err := svc.CreateWebhook(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return updateSnapshot(stack, "CreateWebhook")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestUpdateSnapshot_DeleteBuildBatch(t *testing.T) {
	svc := New(Options{})
	_, err := svc.DeleteBuildBatch(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return updateSnapshot(stack, "DeleteBuildBatch")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestUpdateSnapshot_DeleteFleet(t *testing.T) {
	svc := New(Options{})
	_, err := svc.DeleteFleet(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return updateSnapshot(stack, "DeleteFleet")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestUpdateSnapshot_DeleteProject(t *testing.T) {
	svc := New(Options{})
	_, err := svc.DeleteProject(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return updateSnapshot(stack, "DeleteProject")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestUpdateSnapshot_DeleteReport(t *testing.T) {
	svc := New(Options{})
	_, err := svc.DeleteReport(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return updateSnapshot(stack, "DeleteReport")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestUpdateSnapshot_DeleteReportGroup(t *testing.T) {
	svc := New(Options{})
	_, err := svc.DeleteReportGroup(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return updateSnapshot(stack, "DeleteReportGroup")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestUpdateSnapshot_DeleteResourcePolicy(t *testing.T) {
	svc := New(Options{})
	_, err := svc.DeleteResourcePolicy(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return updateSnapshot(stack, "DeleteResourcePolicy")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestUpdateSnapshot_DeleteSourceCredentials(t *testing.T) {
	svc := New(Options{})
	_, err := svc.DeleteSourceCredentials(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return updateSnapshot(stack, "DeleteSourceCredentials")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestUpdateSnapshot_DeleteWebhook(t *testing.T) {
	svc := New(Options{})
	_, err := svc.DeleteWebhook(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return updateSnapshot(stack, "DeleteWebhook")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestUpdateSnapshot_DescribeCodeCoverages(t *testing.T) {
	svc := New(Options{})
	_, err := svc.DescribeCodeCoverages(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return updateSnapshot(stack, "DescribeCodeCoverages")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestUpdateSnapshot_DescribeTestCases(t *testing.T) {
	svc := New(Options{})
	_, err := svc.DescribeTestCases(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return updateSnapshot(stack, "DescribeTestCases")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestUpdateSnapshot_GetReportGroupTrend(t *testing.T) {
	svc := New(Options{})
	_, err := svc.GetReportGroupTrend(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return updateSnapshot(stack, "GetReportGroupTrend")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestUpdateSnapshot_GetResourcePolicy(t *testing.T) {
	svc := New(Options{})
	_, err := svc.GetResourcePolicy(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return updateSnapshot(stack, "GetResourcePolicy")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestUpdateSnapshot_ImportSourceCredentials(t *testing.T) {
	svc := New(Options{})
	_, err := svc.ImportSourceCredentials(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return updateSnapshot(stack, "ImportSourceCredentials")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestUpdateSnapshot_InvalidateProjectCache(t *testing.T) {
	svc := New(Options{})
	_, err := svc.InvalidateProjectCache(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return updateSnapshot(stack, "InvalidateProjectCache")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestUpdateSnapshot_ListBuildBatches(t *testing.T) {
	svc := New(Options{})
	_, err := svc.ListBuildBatches(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return updateSnapshot(stack, "ListBuildBatches")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestUpdateSnapshot_ListBuildBatchesForProject(t *testing.T) {
	svc := New(Options{})
	_, err := svc.ListBuildBatchesForProject(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return updateSnapshot(stack, "ListBuildBatchesForProject")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestUpdateSnapshot_ListBuilds(t *testing.T) {
	svc := New(Options{})
	_, err := svc.ListBuilds(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return updateSnapshot(stack, "ListBuilds")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestUpdateSnapshot_ListBuildsForProject(t *testing.T) {
	svc := New(Options{})
	_, err := svc.ListBuildsForProject(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return updateSnapshot(stack, "ListBuildsForProject")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestUpdateSnapshot_ListCuratedEnvironmentImages(t *testing.T) {
	svc := New(Options{})
	_, err := svc.ListCuratedEnvironmentImages(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return updateSnapshot(stack, "ListCuratedEnvironmentImages")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestUpdateSnapshot_ListFleets(t *testing.T) {
	svc := New(Options{})
	_, err := svc.ListFleets(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return updateSnapshot(stack, "ListFleets")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestUpdateSnapshot_ListProjects(t *testing.T) {
	svc := New(Options{})
	_, err := svc.ListProjects(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return updateSnapshot(stack, "ListProjects")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestUpdateSnapshot_ListReportGroups(t *testing.T) {
	svc := New(Options{})
	_, err := svc.ListReportGroups(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return updateSnapshot(stack, "ListReportGroups")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestUpdateSnapshot_ListReports(t *testing.T) {
	svc := New(Options{})
	_, err := svc.ListReports(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return updateSnapshot(stack, "ListReports")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestUpdateSnapshot_ListReportsForReportGroup(t *testing.T) {
	svc := New(Options{})
	_, err := svc.ListReportsForReportGroup(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return updateSnapshot(stack, "ListReportsForReportGroup")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestUpdateSnapshot_ListSharedProjects(t *testing.T) {
	svc := New(Options{})
	_, err := svc.ListSharedProjects(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return updateSnapshot(stack, "ListSharedProjects")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestUpdateSnapshot_ListSharedReportGroups(t *testing.T) {
	svc := New(Options{})
	_, err := svc.ListSharedReportGroups(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return updateSnapshot(stack, "ListSharedReportGroups")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestUpdateSnapshot_ListSourceCredentials(t *testing.T) {
	svc := New(Options{})
	_, err := svc.ListSourceCredentials(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return updateSnapshot(stack, "ListSourceCredentials")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestUpdateSnapshot_PutResourcePolicy(t *testing.T) {
	svc := New(Options{})
	_, err := svc.PutResourcePolicy(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return updateSnapshot(stack, "PutResourcePolicy")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestUpdateSnapshot_RetryBuild(t *testing.T) {
	svc := New(Options{})
	_, err := svc.RetryBuild(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return updateSnapshot(stack, "RetryBuild")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestUpdateSnapshot_RetryBuildBatch(t *testing.T) {
	svc := New(Options{})
	_, err := svc.RetryBuildBatch(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return updateSnapshot(stack, "RetryBuildBatch")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestUpdateSnapshot_StartBuild(t *testing.T) {
	svc := New(Options{})
	_, err := svc.StartBuild(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return updateSnapshot(stack, "StartBuild")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestUpdateSnapshot_StartBuildBatch(t *testing.T) {
	svc := New(Options{})
	_, err := svc.StartBuildBatch(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return updateSnapshot(stack, "StartBuildBatch")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestUpdateSnapshot_StopBuild(t *testing.T) {
	svc := New(Options{})
	_, err := svc.StopBuild(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return updateSnapshot(stack, "StopBuild")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestUpdateSnapshot_StopBuildBatch(t *testing.T) {
	svc := New(Options{})
	_, err := svc.StopBuildBatch(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return updateSnapshot(stack, "StopBuildBatch")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestUpdateSnapshot_UpdateFleet(t *testing.T) {
	svc := New(Options{})
	_, err := svc.UpdateFleet(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return updateSnapshot(stack, "UpdateFleet")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestUpdateSnapshot_UpdateProject(t *testing.T) {
	svc := New(Options{})
	_, err := svc.UpdateProject(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return updateSnapshot(stack, "UpdateProject")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestUpdateSnapshot_UpdateProjectVisibility(t *testing.T) {
	svc := New(Options{})
	_, err := svc.UpdateProjectVisibility(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return updateSnapshot(stack, "UpdateProjectVisibility")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestUpdateSnapshot_UpdateReportGroup(t *testing.T) {
	svc := New(Options{})
	_, err := svc.UpdateReportGroup(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return updateSnapshot(stack, "UpdateReportGroup")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestUpdateSnapshot_UpdateWebhook(t *testing.T) {
	svc := New(Options{})
	_, err := svc.UpdateWebhook(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return updateSnapshot(stack, "UpdateWebhook")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

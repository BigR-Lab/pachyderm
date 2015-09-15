package pretty

import (
	"fmt"
	"io"
	"time"

	"go.pedge.io/proto/time"

	"github.com/docker/docker/pkg/units"
	"github.com/pachyderm/pachyderm/src/pfs"
)

func PrintRepoHeader(w io.Writer) {
	fmt.Fprintln(w, "NAME\tCREATED\tSIZE\t")
}

func PrintRepoInfo(w io.Writer, repoInfo *pfs.RepoInfo) {
	fmt.Fprintf(w, "%s\t", repoInfo.Repo.Name)
	fmt.Fprintf(
		w,
		"%s ago\t", units.HumanDuration(
			time.Since(
				prototime.TimestampToTime(
					repoInfo.Created,
				),
			),
		),
	)
	fmt.Fprintf(w, "%s\t\n", units.BytesSize(float64(repoInfo.SizeBytes)))
}

func PrintCommitInfoHeader(w io.Writer) {
	fmt.Fprintln(w, "ID\tPARENT\tSTATUS\tTIME_OPENED\tTIME_CLOSED\tTOTAL_SIZE\tDIFF_SIZE\t")
}

func PrintCommitInfo(w io.Writer, commitInfo *pfs.CommitInfo) {
	fmt.Fprintf(w, "%s\t", commitInfo.Commit.Id)
	if commitInfo.ParentCommit != nil {
		fmt.Fprintf(w, "%s\t", commitInfo.ParentCommit.Id)
	} else {
		fmt.Fprint(w, "<none>\t")
	}
	if commitInfo.CommitType == pfs.CommitType_COMMIT_TYPE_WRITE {
		fmt.Fprint(w, "writeable\t")
	} else {
		fmt.Fprint(w, "read-only\t")
	}
	fmt.Fprint(w, "-\t")
	fmt.Fprint(w, "-\t")
	fmt.Fprint(w, "-\t")
	fmt.Fprint(w, "-\t\n")
}

func PrintFileInfoHeader(w io.Writer) {
	fmt.Fprintln(w, "NAME\tTYPE\tMODIFIED\tLAST_COMMIT_MODIFIED\tSIZE\tPERMISSIONS\t")
}

func PrintFileInfo(w io.Writer, fileInfo *pfs.FileInfo) {
	fmt.Fprintf(w, "%s\t", fileInfo.File.Path)
	if fileInfo.FileType == pfs.FileType_FILE_TYPE_REGULAR {
		fmt.Fprint(w, "file\t")
	} else {
		fmt.Fprint(w, "dir\t")
	}
	fmt.Fprintf(
		w,
		"%s ago\t", units.HumanDuration(
			time.Since(
				prototime.TimestampToTime(
					fileInfo.LastModified,
				),
			),
		),
	)
	fmt.Fprint(w, "-\t")
	fmt.Fprintf(w, "%s\t", units.BytesSize(float64(fileInfo.SizeBytes)))
	fmt.Fprintf(w, "%4d\t\n", fileInfo.Perm)
}
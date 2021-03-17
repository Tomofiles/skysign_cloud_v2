package action

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TelemetrySnapshotをTrajectoryに追加し、軌道を伸長する。
func TestExtension(t *testing.T) {
	a := assert.New(t)

	preTrajectory := Trajectory{
		numberOfPoints:   0,
		trajectoryPoints: []trajectoryPoint{},
	}

	postTrajectory := preTrajectory.Extension(DefaultTelemetrySnapshot)

	a.Equal(preTrajectory.numberOfPoints, 0)
	a.Len(preTrajectory.trajectoryPoints, 0)
	a.Equal(postTrajectory.numberOfPoints, 1)
	a.Len(postTrajectory.trajectoryPoints, 1)
	a.Equal(postTrajectory.trajectoryPoints[0], DefaultTrajectoryPoint)
}

// TrajectoryのSnapshotを取得する。
// Trajectoryが複数件の場合、返却されるSnapshotリストが
// 同一件数であることを検証する。
func TestTakeASnapshots_MultipleTrajectoryPoint(t *testing.T) {
	a := assert.New(t)

	trajectory := Trajectory{
		numberOfPoints: 3,
		trajectoryPoints: []trajectoryPoint{
			DefaultTrajectoryPoint,
			DefaultTrajectoryPoint,
			DefaultTrajectoryPoint,
		},
	}

	snapshots := trajectory.TakeASnapshots()

	a.Len(snapshots, 3)
	a.Equal(snapshots[0], DefaultTelemetrySnapshot)
	a.Equal(snapshots[1], DefaultTelemetrySnapshot)
	a.Equal(snapshots[2], DefaultTelemetrySnapshot)
}

// TrajectoryのSnapshotを取得する。
// Trajectoryが0件の場合、返却されるSnapshotリストが
// 0件であることを検証する。
func TestTakeASnapshots_NoTrajectoryPoint(t *testing.T) {
	a := assert.New(t)

	trajectory := Trajectory{
		numberOfPoints:   0,
		trajectoryPoints: []trajectoryPoint{},
	}

	snapshots := trajectory.TakeASnapshots()

	a.Len(snapshots, 0)
}

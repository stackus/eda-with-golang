package es

type SnapshotPayload interface {
	SnapshotName() string
}

type Snapshot interface {
	Version() int
	Payload() SnapshotPayload
}

type snapshot struct {
	payload SnapshotPayload
	version int
}

type Snapshotted interface {
	ToSnapshot() (SnapshotPayload, error)
	LoadSnapshot(snapshot Snapshot) error
}

func NewSnapshot(payload SnapshotPayload, version int) *snapshot {
	return &snapshot{
		payload: payload,
		version: version,
	}
}

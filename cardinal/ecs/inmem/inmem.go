// Package inmem is a helper package that allows for the creation of an *ecs.World object
// that uses an in-memory redis DB as the storage layer. This is useful for local development
// or for tests. Data will not be persisted between runs, so this is not suitable for any
// kind of prodcution or staging environemnts.
package inmem

import (
	"testing"

	"github.com/alicebob/miniredis/v2"
	"github.com/argus-labs/world-engine/cardinal/ecs"
	"github.com/argus-labs/world-engine/cardinal/ecs/storage"
)

// NewECSWorld creates an ecs.World that uses an in-memory redis DB as the storage
// layer. This is only suitable for local development. If you are creating an ecs.World for
// unit tests, use NewECSWorldForTest.
func NewECSWorld() *ecs.World {
	s, err := miniredis.Run()
	if err != nil {
		panic("Unable to initialize in-memory redis")
	}
	return newInMemoryWorld(s)
}

// NewECSWorldForTest creates an ecs.World suitable for running in tests. Relevant resources
// are automatically cleaned up at the completion of each test.
func NewECSWorldForTest(t testing.TB) *ecs.World {
	s := miniredis.RunT(t)
	return newInMemoryWorld(s)
}

func newInMemoryWorld(s *miniredis.Miniredis) *ecs.World {
	rs := storage.NewRedisStorage(storage.Options{
		Addr:     s.Addr(),
		Password: "", // no password set
		DB:       0,  // use default DB
	}, "in-memory-world")
	worldStorage := storage.NewWorldStorage(
		storage.Components{Store: &rs, ComponentIndices: &rs},
		&rs,
		storage.NewArchetypeComponentIndex(),
		storage.NewArchetypeAccessor(),
		&rs,
		&rs)

	return ecs.NewWorld(worldStorage)
}

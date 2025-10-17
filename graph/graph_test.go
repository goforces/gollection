package graph

import (
	"sort"
	"testing"
)

func TestNew(t *testing.T) {
	t.Run("directed", func(t *testing.T) {
		g := New[int](true)
		if g == nil {
			t.Fatal("New() returned nil")
		}
		if !g.directed {
			t.Error("graph should be directed")
		}
	})
	
	t.Run("undirected", func(t *testing.T) {
		g := New[int](false)
		if g == nil {
			t.Fatal("New() returned nil")
		}
		if g.directed {
			t.Error("graph should be undirected")
		}
	})
}

func TestAddVertex(t *testing.T) {
	g := New[int](true)
	
	g.AddVertex(1)
	g.AddVertex(2)
	g.AddVertex(3)
	
	vertices := g.Vertices()
	if len(vertices) != 3 {
		t.Errorf("vertices count = %d, want 3", len(vertices))
	}
	
	// Adding duplicate should not increase count
	g.AddVertex(1)
	vertices = g.Vertices()
	if len(vertices) != 3 {
		t.Errorf("vertices count after duplicate = %d, want 3", len(vertices))
	}
}

func TestAddEdgeDirected(t *testing.T) {
	g := New[int](true)
	
	g.AddEdge(1, 2, 1.0)
	
	// Verify edge exists 1->2
	neighbors := g.Neighbors(1)
	if weight, ok := neighbors[2]; !ok || weight != 1.0 {
		t.Errorf("edge 1->2 not found or wrong weight")
	}
	
	// Verify reverse edge doesn't exist (directed)
	neighbors = g.Neighbors(2)
	if _, ok := neighbors[1]; ok {
		t.Error("reverse edge 2->1 should not exist in directed graph")
	}
}

func TestAddEdgeUndirected(t *testing.T) {
	g := New[int](false)
	
	g.AddEdge(1, 2, 1.0)
	
	// Verify edge exists 1->2
	neighbors := g.Neighbors(1)
	if weight, ok := neighbors[2]; !ok || weight != 1.0 {
		t.Errorf("edge 1->2 not found or wrong weight")
	}
	
	// Verify reverse edge exists (undirected)
	neighbors = g.Neighbors(2)
	if weight, ok := neighbors[1]; !ok || weight != 1.0 {
		t.Errorf("edge 2->1 not found or wrong weight in undirected graph")
	}
}

func TestNeighbors(t *testing.T) {
	g := New[int](true)
	
	g.AddEdge(1, 2, 1.0)
	g.AddEdge(1, 3, 2.0)
	g.AddEdge(1, 4, 3.0)
	
	neighbors := g.Neighbors(1)
	if len(neighbors) != 3 {
		t.Errorf("neighbor count = %d, want 3", len(neighbors))
	}
	
	expected := map[int]float64{2: 1.0, 3: 2.0, 4: 3.0}
	for v, w := range expected {
		if neighbors[v] != w {
			t.Errorf("neighbor[%d] = %f, want %f", v, neighbors[v], w)
		}
	}
}

func TestNeighborsEmpty(t *testing.T) {
	g := New[int](true)
	
	g.AddVertex(1)
	
	neighbors := g.Neighbors(1)
	if len(neighbors) != 0 {
		t.Errorf("vertex with no edges should have 0 neighbors, got %d", len(neighbors))
	}
}

func TestNeighborsNonExistent(t *testing.T) {
	g := New[int](true)
	
	neighbors := g.Neighbors(999)
	if neighbors != nil {
		t.Error("neighbors of non-existent vertex should be nil")
	}
}

func TestVertices(t *testing.T) {
	g := New[string](true)
	
	g.AddVertex("a")
	g.AddVertex("b")
	g.AddVertex("c")
	
	vertices := g.Vertices()
	if len(vertices) != 3 {
		t.Fatalf("vertices count = %d, want 3", len(vertices))
	}
	
	// Sort for comparison
	sort.Strings(vertices)
	want := []string{"a", "b", "c"}
	
	for i := range vertices {
		if vertices[i] != want[i] {
			t.Errorf("vertex[%d] = %q, want %q", i, vertices[i], want[i])
		}
	}
}

func TestRemoveVertex(t *testing.T) {
	g := New[int](true)
	
	// Build graph: 1->2, 2->3, 3->1
	g.AddEdge(1, 2, 1.0)
	g.AddEdge(2, 3, 1.0)
	g.AddEdge(3, 1, 1.0)
	
	// Remove vertex 2
	g.RemoveVertex(2)
	
	// Verify vertex 2 is gone
	vertices := g.Vertices()
	for _, v := range vertices {
		if v == 2 {
			t.Error("vertex 2 should be removed")
		}
	}
	
	// Verify edge 1->2 is gone
	neighbors := g.Neighbors(1)
	if _, ok := neighbors[2]; ok {
		t.Error("edge to removed vertex should be gone")
	}
	
	// Verify edge 3->2 is gone (incoming edges removed)
	neighbors = g.Neighbors(3)
	if _, ok := neighbors[2]; ok {
		t.Error("edge to removed vertex should be gone")
	}
	
	// Verify edge 3->1 still exists
	neighbors = g.Neighbors(3)
	if _, ok := neighbors[1]; !ok {
		t.Error("unrelated edge should still exist")
	}
}

func TestRemoveVertexUndirected(t *testing.T) {
	g := New[int](false)
	
	g.AddEdge(1, 2, 1.0)
	g.AddEdge(2, 3, 1.0)
	
	g.RemoveVertex(2)
	
	// Both edges involving 2 should be gone
	neighbors1 := g.Neighbors(1)
	if _, ok := neighbors1[2]; ok {
		t.Error("edge 1-2 should be removed")
	}
	
	neighbors3 := g.Neighbors(3)
	if _, ok := neighbors3[2]; ok {
		t.Error("edge 3-2 should be removed")
	}
}

func TestRemoveEdgeDirected(t *testing.T) {
	g := New[int](true)
	
	g.AddEdge(1, 2, 1.0)
	g.AddEdge(1, 3, 1.0)
	
	g.RemoveEdge(1, 2)
	
	// Edge 1->2 should be gone
	neighbors := g.Neighbors(1)
	if _, ok := neighbors[2]; ok {
		t.Error("removed edge should not exist")
	}
	
	// Edge 1->3 should still exist
	if _, ok := neighbors[3]; !ok {
		t.Error("unrelated edge should still exist")
	}
}

func TestRemoveEdgeUndirected(t *testing.T) {
	g := New[int](false)
	
	g.AddEdge(1, 2, 1.0)
	
	g.RemoveEdge(1, 2)
	
	// Both directions should be removed
	neighbors1 := g.Neighbors(1)
	if _, ok := neighbors1[2]; ok {
		t.Error("edge 1->2 should be removed")
	}
	
	neighbors2 := g.Neighbors(2)
	if _, ok := neighbors2[1]; ok {
		t.Error("edge 2->1 should be removed")
	}
}

func TestRemoveEdgeNonExistent(t *testing.T) {
	g := New[int](true)
	
	g.AddVertex(1)
	g.AddVertex(2)
	
	// Should not panic
	g.RemoveEdge(1, 2)
	g.RemoveEdge(999, 1000)
}

func TestClone(t *testing.T) {
	original := New[int](true)
	
	original.AddEdge(1, 2, 1.0)
	original.AddEdge(2, 3, 2.0)
	original.AddEdge(3, 1, 3.0)
	
	clone := original.Clone()
	
	// Verify directed flag copied
	if clone.directed != original.directed {
		t.Error("clone should have same directed flag")
	}
	
	// Verify vertices match
	origVerts := original.Vertices()
	cloneVerts := clone.Vertices()
	if len(origVerts) != len(cloneVerts) {
		t.Errorf("vertex count: original=%d, clone=%d", len(origVerts), len(cloneVerts))
	}
	
	// Verify edges match
	for _, v := range origVerts {
		origNeighbors := original.Neighbors(v)
		cloneNeighbors := clone.Neighbors(v)
		
		if len(origNeighbors) != len(cloneNeighbors) {
			t.Errorf("vertex %d neighbor count mismatch", v)
		}
		
		for neighbor, weight := range origNeighbors {
			if cloneNeighbors[neighbor] != weight {
				t.Errorf("edge %d->%d weight mismatch", v, neighbor)
			}
		}
	}
}

func TestCloneIndependence(t *testing.T) {
	original := New[int](true)
	
	original.AddEdge(1, 2, 1.0)
	
	clone := original.Clone()
	
	// Modify clone
	clone.AddEdge(2, 3, 2.0)
	
	// Original should be unaffected
	neighbors := original.Neighbors(2)
	if _, ok := neighbors[3]; ok {
		t.Error("modifying clone should not affect original")
	}
}

func TestWeights(t *testing.T) {
	g := New[int](true)
	
	g.AddEdge(1, 2, 5.5)
	g.AddEdge(1, 3, 10.7)
	g.AddEdge(2, 3, 3.14)
	
	tests := []struct {
		from, to int
		want     float64
	}{
		{1, 2, 5.5},
		{1, 3, 10.7},
		{2, 3, 3.14},
	}
	
	for _, tt := range tests {
		neighbors := g.Neighbors(tt.from)
		if weight := neighbors[tt.to]; weight != tt.want {
			t.Errorf("edge %d->%d weight = %f, want %f", tt.from, tt.to, weight, tt.want)
		}
	}
}

func TestNegativeWeights(t *testing.T) {
	g := New[int](true)
	
	g.AddEdge(1, 2, -5.0)
	
	neighbors := g.Neighbors(1)
	if weight := neighbors[2]; weight != -5.0 {
		t.Errorf("negative weight = %f, want -5.0", weight)
	}
}

func TestZeroWeight(t *testing.T) {
	g := New[int](true)
	
	g.AddEdge(1, 2, 0.0)
	
	neighbors := g.Neighbors(1)
	if weight, ok := neighbors[2]; !ok || weight != 0.0 {
		t.Error("edge with zero weight should exist")
	}
}

func TestStringVertices(t *testing.T) {
	g := New[string](true)
	
	g.AddEdge("alice", "bob", 1.0)
	g.AddEdge("bob", "charlie", 2.0)
	g.AddEdge("charlie", "alice", 3.0)
	
	neighbors := g.Neighbors("alice")
	if weight, ok := neighbors["bob"]; !ok || weight != 1.0 {
		t.Error("graph should work with string vertices")
	}
}

func TestComplexGraph(t *testing.T) {
	g := New[int](true)
	
	// Build a more complex graph
	edges := []struct {
		from, to int
		weight   float64
	}{
		{1, 2, 1.0},
		{1, 3, 4.0},
		{2, 3, 2.0},
		{2, 4, 5.0},
		{3, 4, 1.0},
		{4, 5, 3.0},
	}
	
	for _, e := range edges {
		g.AddEdge(e.from, e.to, e.weight)
	}
	
	// Verify all edges exist
	for _, e := range edges {
		neighbors := g.Neighbors(e.from)
		if weight, ok := neighbors[e.to]; !ok || weight != e.weight {
			t.Errorf("edge %d->%d missing or wrong weight", e.from, e.to)
		}
	}
	
	// Verify vertex count
	vertices := g.Vertices()
	if len(vertices) != 5 {
		t.Errorf("vertex count = %d, want 5", len(vertices))
	}
}

func TestSelfLoop(t *testing.T) {
	g := New[int](true)
	
	g.AddEdge(1, 1, 1.0)
	
	neighbors := g.Neighbors(1)
	if weight, ok := neighbors[1]; !ok || weight != 1.0 {
		t.Error("self-loop should be allowed")
	}
}

func TestMultipleEdgesSameVertices(t *testing.T) {
	g := New[int](true)
	
	// Add edge, then update weight
	g.AddEdge(1, 2, 1.0)
	g.AddEdge(1, 2, 2.0)
	
	neighbors := g.Neighbors(1)
	if weight := neighbors[2]; weight != 2.0 {
		t.Errorf("edge weight = %f, want 2.0 (should update)", weight)
	}
}

func TestIsolatedVertex(t *testing.T) {
	g := New[int](true)
	
	g.AddVertex(1)
	g.AddEdge(2, 3, 1.0)
	
	// Vertex 1 has no edges
	neighbors := g.Neighbors(1)
	if len(neighbors) != 0 {
		t.Error("isolated vertex should have no neighbors")
	}
}

func TestEmptyGraph(t *testing.T) {
	g := New[int](true)
	
	vertices := g.Vertices()
	if len(vertices) != 0 {
		t.Errorf("empty graph should have 0 vertices, got %d", len(vertices))
	}
}

func TestLargeGraph(t *testing.T) {
	g := New[int](true)
	n := 1000
	
	// Add vertices
	for i := 0; i < n; i++ {
		g.AddVertex(i)
	}
	
	// Add edges (each vertex connected to next)
	for i := 0; i < n-1; i++ {
		g.AddEdge(i, i+1, float64(i))
	}
	
	// Verify vertex count
	vertices := g.Vertices()
	if len(vertices) != n {
		t.Errorf("vertex count = %d, want %d", len(vertices), n)
	}
	
	// Spot check some edges
	for i := 0; i < 10; i++ {
		neighbors := g.Neighbors(i)
		if weight, ok := neighbors[i+1]; !ok || weight != float64(i) {
			t.Errorf("edge %d->%d missing or wrong weight", i, i+1)
		}
	}
}

func TestUndirectedSymmetry(t *testing.T) {
	g := New[int](false)
	
	g.AddEdge(1, 2, 5.0)
	
	// Both directions should have same weight
	neighbors1 := g.Neighbors(1)
	neighbors2 := g.Neighbors(2)
	
	if neighbors1[2] != neighbors2[1] {
		t.Error("undirected edges should have symmetric weights")
	}
}

func TestCompleteRemoval(t *testing.T) {
	g := New[int](true)
	
	// Build graph
	g.AddEdge(1, 2, 1.0)
	g.AddEdge(2, 3, 1.0)
	g.AddEdge(3, 1, 1.0)
	
	// Remove all vertices
	g.RemoveVertex(1)
	g.RemoveVertex(2)
	g.RemoveVertex(3)
	
	vertices := g.Vertices()
	if len(vertices) != 0 {
		t.Errorf("after removing all vertices, count = %d, want 0", len(vertices))
	}
}

// Benchmarks
func BenchmarkAddVertex(b *testing.B) {
	g := New[int](true)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		g.AddVertex(i)
	}
}

func BenchmarkAddEdge(b *testing.B) {
	g := New[int](true)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		g.AddEdge(i, i+1, 1.0)
	}
}

func BenchmarkNeighbors(b *testing.B) {
	g := New[int](true)
	for i := 0; i < 1000; i++ {
		g.AddEdge(0, i, 1.0)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		g.Neighbors(0)
	}
}

func BenchmarkRemoveVertex(b *testing.B) {
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		g := New[int](true)
		for j := 0; j < 1000; j++ {
			g.AddVertex(j)
			if j > 0 {
				g.AddEdge(j-1, j, 1.0)
			}
		}
		b.StartTimer()
		g.RemoveVertex(500)
		b.StopTimer()
	}
}


package graph

// Graph is a simple adjacency-list graph supporting directed or undirected edges.
type Graph[T comparable] struct {
	directed bool
	adj      map[T]map[T]float64
}

// New creates a new graph. If directed is true, edges are one-way.
func New[T comparable](directed bool) *Graph[T] {
	return &Graph[T]{directed: directed, adj: make(map[T]map[T]float64)}
}

// AddVertex ensures the vertex exists.
func (g *Graph[T]) AddVertex(v T) {
	if _, ok := g.adj[v]; !ok {
		g.adj[v] = make(map[T]float64)
	}
}

// AddEdge adds an edge u->v with weight w. For undirected graphs, adds both ways.
func (g *Graph[T]) AddEdge(u, v T, w float64) {
	g.AddVertex(u)
	g.AddVertex(v)
	g.adj[u][v] = w
	if !g.directed {
		g.adj[v][u] = w
	}
}

// Neighbors returns the neighbor-weight map for v (may be empty).
func (g *Graph[T]) Neighbors(v T) map[T]float64 { return g.adj[v] }

// Vertices returns all vertices.
func (g *Graph[T]) Vertices() []T {
	out := make([]T, 0, len(g.adj))
	for v := range g.adj {
		out = append(out, v)
	}
	return out
}

// RemoveVertex removes a vertex and all edges connected to it.
func (g *Graph[T]) RemoveVertex(v T) {
	// Remove all edges pointing to v from other vertices
	for vertex := range g.adj {
		delete(g.adj[vertex], v)
	}
	// Remove the vertex itself
	delete(g.adj, v)
}

// RemoveEdge removes an edge from u to v.
// For undirected graphs, removes both u->v and v->u.
func (g *Graph[T]) RemoveEdge(u, v T) {
	if neighbors, ok := g.adj[u]; ok {
		delete(neighbors, v)
	}
	if !g.directed {
		if neighbors, ok := g.adj[v]; ok {
			delete(neighbors, u)
		}
	}
}

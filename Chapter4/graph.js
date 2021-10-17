export class Node {
    constructor(value) {
      this.value = value;
      this.adjacents = []; // adjacency list
    }
  
    addAdjacent(node) {
      this.adjacents.push(node);
    }
  
    removeAdjacent(node) {
      const index = this.adjacents.indexOf(node);
      if(index > -1) {
        this.adjacents.splice(index, 1);
        return node;
      }
    }
  
    getAdjacents() {
      return this.adjacents;
    }
  
    isAdjacent(node) {
      return this.adjacents.indexOf(node) > -1;
    }
}

export default class Graph {
    constructor(edgeDirection = Graph.DIRECTED) {
      this.nodes = new Map();
      this.edgeDirection = edgeDirection;
    }

    addEdge(source, destination) {
        const sourceNode = this.addVertex(source);
        const destinationNode = this.addVertex(destination);
      
        sourceNode.addAdjacent(destinationNode);
      
        if(this.edgeDirection === Graph.UNDIRECTED) {
          destinationNode.addAdjacent(sourceNode);
        }
      
        return [sourceNode, destinationNode];
      }

      addVertex(value) {
        if(this.nodes.has(value)) {
          return this.nodes.get(value);
        } else {
          const vertex = new Node(value);
          this.nodes.set(value, vertex);
          return vertex;
        }
      }      
    
      removeVertex(value) {
        const current = this.nodes.get(value);
        if(current) {
          for (const node of this.nodes.values()) {
            node.removeAdjacent(current);
          }
        }
        return this.nodes.delete(value);
      }

      removeEdge(source, destination) {
        const sourceNode = this.nodes.get(source);
        const destinationNode = this.nodes.get(destination);
      
        if(sourceNode && destinationNode) {
          sourceNode.removeAdjacent(destinationNode);
      
          if(this.edgeDirection === Graph.UNDIRECTED) {
            destinationNode.removeAdjacent(sourceNode);
          }
        }
      
        return [sourceNode, destinationNode];
      }
}


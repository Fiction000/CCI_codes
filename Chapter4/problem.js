import Graph from './graph.js';

function checkPath(node1, node2) {
    const visited = new Map();
    const visitList = [];

    visitList.push(node1);

    while(visitList.length !== 0) {
        const node = visitList.pop();
        if(node && !visited.has(node)) {
            visited.set(node);
            for (let adj of node.getAdjacents()) {
                if (adj === node2) {
                    return true;
                }
                visitList.push(adj)
            }   
        }
    }
    return false;
}

const graph = new Graph(Graph.DIRECTED);

const [first] = graph.addEdge(1, 2);
graph.addEdge(3, 4)
graph.addEdge(2, 3)


const node1 = graph.nodes.get(1);
const node2 = graph.nodes.get(3);
const result = checkPath(node1, node2);

console.log(result);
// result = checkPath(graph[1], graph[2]);
// console.log(result);

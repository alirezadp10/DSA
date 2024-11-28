package huffman

import "fmt"

func decode(s string, m map[string]byte) string {
    res := ""
    k := ""
    for _, v := range s {
        k += string(v)
        if val, ok := m[k]; ok {
            res += string(val)
            k = ""
        }
    }
    return res
}

func encode(s string) string {
    // enumerating
    m := make(map[byte]int)
    for _, i := range s {
        m[byte(i)]++
    }

    // making nodes
    var nl []Node
    for l, v := range m {
        nl = append(nl, Node{value: v, label: l})
    }

    // making graph
    graph := makeGraph(nl)

    // mapping keys
    k := map[byte]string{}
    graphToMap(&graph, k, "")

    // making the result
    res := ""
    for _, c := range s {
        res += k[byte(c)]
    }

    fmt.Printf("%v", k)

    return res
}

func sort(nl []Node) []Node {
    if len(nl) <= 1 {
        return nl
    }
    p := len(nl) - 1

    l := 0
    r := p - 1

    for l != r {
        if nl[l].value < nl[p].value {
            l++
            continue
        }
        if nl[r].value > nl[p].value {
            r--
            continue
        }
        temp := nl[l]
        nl[l] = nl[r]
        nl[r] = temp
    }

    if l != p-1 {
        temp := nl[l]
        nl[l] = nl[p]
        nl[p] = temp
    }

    return append(sort(nl[:l+1]), sort(nl[l+1:])...)
}

func makeGraph(nl []Node) Node {
    if len(nl) == 1 {
        return nl[0]
    }

    nl = sort(nl)

    nl = append(nl[2:], Node{
        value: nl[0].value + nl[1].value,
        left:  &nl[0],
        right: &nl[1],
    })

    return makeGraph(nl)
}

func graphToMap(n *Node, m map[byte]string, val string) {
    if n.right == nil && n.left == nil {
        m[n.label] = val
    }

    if n.left != nil {
        graphToMap(n.left, m, val+"0")
    }
    if n.right != nil {
        graphToMap(n.right, m, val+"1")
    }
}

type Node struct {
    left  *Node
    right *Node
    value int
    label byte
}

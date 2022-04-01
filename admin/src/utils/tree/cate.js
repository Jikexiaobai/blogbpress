export function loopCate(tree){
    return tree.map(item =>{
    const newNode = {
        key: item.cateId,
        title: item.title,
        value: item.cateId,
    }
    
        // 是否有子菜单，并递归处理
    if (item.children && item.children.length > 0) {
    // Recursion
        newNode.children = loopCate(item.children)
    }
        return newNode
    })
}
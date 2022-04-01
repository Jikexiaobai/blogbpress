// 菜单树遍历
export function loopMenu(tree,pid = 0){
    var itemArr = [];
    tree.forEach(item => {
        // 判断是否为父级菜单
        if (item.parentId === pid) {
            const child = {
                ...item,
                key: item.id,
                value: item.id,
                children: []
            }
            child.children = loopMenu(tree, item.id)
            itemArr.push(child)
        }
    })
    return itemArr;
}
  
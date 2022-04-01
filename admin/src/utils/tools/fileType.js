  // 获取文件后缀
export function getType(file){
    var filename=file;
    var index1=filename.lastIndexOf(".");
    var index2=filename.length;
    var type=filename.substring(index1,index2);
    return type;
}

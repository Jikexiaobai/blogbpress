
import { differenceInHours,differenceInDays,differenceInMonths,differenceInMinutes } from "date-fns"


// 日期修改
export function resetData(date){
    let str = date
   
    date = date.replace(/-/g, '/');
    const startDate = new Date(date);
    const endDate = new Date();

    let inMinutes = differenceInMinutes(endDate, startDate)
    if (inMinutes > 0 && inMinutes < 10) {
        str = `刚刚`
        return str
    }
    
    if (inMinutes >= 10 && inMinutes < 60) {
        str = `${inMinutes}分钟之前`
        return str
    }

    const inHours = differenceInHours(endDate, startDate)
    if (inHours >= 1 && inHours < 24) {
        str = `${inHours}小时之前`
        return str
    }

    const inDays = differenceInDays(endDate, startDate)
    if (inDays >= 1 && inDays < 31) {
        str = `${inDays}天之前`
        return str
    }

    const inMonths = differenceInMonths(endDate, startDate)
    if (inMonths >= 1 && inMonths < 12) {
        str = `${inMonths}月之前`
        return str
    }
    
    return str;
}

// 日期修改
export function resetNum(count){
    let str = ""
   
    if (count >= 0 && count < 999) {
        str = `${count}`
        return str
    }
    
    if (count >= 999 && count <= 9999) {
        str = `${Math.round(count / 1000)}k`
        return str
    }

    if (count >= 9999 && count <= 99999) {
        str = `${Math.round(count / 10000)}w`
        return str
    }

    if (count > 99999) {
        str = `${10}w+`
        return str
    }
    return str;
}

// 缩略图
export function resetImage(v,w = null,h = null){
    if (w == null || h == null) { 
        return v;
    }
    v = v+`@w${w}_h${h}`
    return v;
}

// deepCopy 简单的deep copy，没有考虑对象的属性还是对象的情况
export function deepCopy<T extends object>(obj: T): T {
    let res: T = {} as T;

    for (let key in obj) {
        res[key] = obj[key];
    }

    return res;
}

// objectToFormData 泛型用于解决`obj[key]`报错问题
export function objectToFormData<T extends object>(obj: T): FormData {
    let data = new FormData()
    for (let key in obj) {
        if (typeof obj[key] == "object") { // if field type is another object
            objectToFormData(obj[key] as object).forEach((value:FormDataEntryValue, key: string) => {
                data.append(key, value)
            })
        } else { // normal
            data.append(key, obj[key] as string)
        }

    }

    return data
}

import { ElMessage } from "element-plus";

// Log print to console and display in webpage
class Log {
    success(text: string): void {
        console.log("> " + text + ".");
        ElMessage({
            message: text,
            type: "success",
        });
    }

    fail(text: string, err: any = ""): void {
        console.log("> " + text + ".", err);
        ElMessage({
            message: text,
            type: "error",
        });
    }
}

export const log = new Log();

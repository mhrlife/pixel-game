export function getInitData(): string {
    let initData;
    try {
         initData = window.Telegram.WebApp.initData;
    } catch (e) {}

    if(initData) {
        return initData;
    }

    return "TEST_TOKEN"
}
export function getInitData(): string {

    return "query_id=AAH7EtEFAAAAAPsS0QWu3hNX&user=%7B%22id%22%3A97587963%2C%22first_name%22%3A%22Mohammad%22%2C%22last_name%22%3A%22Hoseini%20Rad%22%2C%22username%22%3A%22pp2007ws%22%2C%22language_code%22%3A%22en%22%2C%22is_premium%22%3Atrue%2C%22allows_write_to_pm%22%3Atrue%2C%22photo_url%22%3A%22https%3A%5C%2F%5C%2Ft.me%5C%2Fi%5C%2Fuserpic%5C%2F320%5C%2FSE2gzUqVYoRQE3n1hbH1KWty_L5SNkYo93CdiDdEwCc.svg%22%7D&auth_date=1732879194&signature=DOJ_HOGOGsr2UJ3BROfWlmnkU8xs86qANeNhnkWLL1ykFKDcmnG9FgTtEekdd-n2bQtjbRtPP9RlD-IzuSEOCg&hash=6f00a943c7b31294f6f362666d96c827c1746069eed297d77f567ccb37288596";

    try {
        return window.Telegram.WebApp.initData
    } catch (e) {
        return ""
    }
}
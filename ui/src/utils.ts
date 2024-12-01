export const forceFarsiNumbers = (str: string | number): string => {
    if (typeof str === 'number') {
        str = str.toString();
    }

    return str.replace(/[0-9]/g, function (d) {
        return String.fromCharCode(d.charCodeAt(0) + 1728);
    });
}
export const Colors = ['red-light', 'red-dark', 'blue-light', 'blue-dark', 'green-light', 'green-dark',
    'yellow-light', 'yellow-dark', 'purple-light', 'purple-dark', 'orange-light', 'orange-dark', 'pink-light',
    'pink-dark', 'cyan-light', 'cyan-dark', 'teal-light', 'teal-dark', 'white', 'black','gray'] as const;

export type Color = typeof Colors[number];


export function colorToHex(color: Color): string {
    switch (color) {
        case 'red-light':
            return '#FFCDD2';
        case 'red-dark':
            return '#EF9A9A';
        case 'blue-light':
            return '#BBDEFB';
        case 'blue-dark':
            return '#64B5F6';
        case 'green-light':
            return '#C8E6C9';
        case 'green-dark':
            return '#81C784';
        case 'yellow-light':
            return '#FFF9C4';
        case 'yellow-dark':
            return '#FFF176';
        case 'purple-light':
            return '#E1BEE7';
        case 'purple-dark':
            return '#BA68C8';
        case 'orange-light':
            return '#FFE0B2';
        case 'orange-dark':
            return '#FFB74D';
        case 'pink-light':
            return '#F8BBD0';
        case 'pink-dark':
            return '#F06292';
        case 'cyan-light':
            return '#B2EBF2';
        case 'cyan-dark':
            return '#4DD0E1';
        case 'teal-light':
            return '#B2DFDB';
        case 'teal-dark':
            return '#4DB6AC';
        case 'white':
            return '#FFFFFF';
        case 'black':
            return '#4c4c4c';
        case 'gray':
            return '#B0BEC5';
    }

    return '#FFFFFF';
}


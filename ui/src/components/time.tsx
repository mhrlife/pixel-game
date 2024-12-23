import {format, formatDistanceToNow} from 'date-fns';

export const TimestampComponent = ({unixTimestamp}: { unixTimestamp: number }) => {
    // Multiply by 1000 if your timestamp is in seconds
    const date = new Date(unixTimestamp * 1000);
    const formattedDate = format(date, 'PPPpp'); // e.g., Jan 1, 2022, 12:00 PM

    return <div>{formattedDate}</div>;
};


export const TimeAgo = ({unixTimestamp}: { unixTimestamp: number }) => {
    const date = new Date(unixTimestamp * 1000);
    const timeAgo = formatDistanceToNow(date, {addSuffix: true}); // e.g., "3 minutes ago"

    return <div>{timeAgo}</div>;
};


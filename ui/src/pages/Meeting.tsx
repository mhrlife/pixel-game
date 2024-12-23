import {useNavigate, useParams} from "react-router";
import {InfinitySpin, ThreeDots} from "react-loader-spinner";
import {useRoomToken} from "@/api/room.ts";
import {
    RiArrowGoBackFill,
    RiEmotionSadLine,
    RiMore2Fill,
    RiShare2Fill,
    RiUser2Line,
    RiVideoLine,
    RiVideoOffLine,
    RiVolumeMuteLine,
    RiVolumeUpLine
} from "react-icons/ri";
import {Button} from "@/components/ui/button.tsx";
import {
    LiveKitRoom,
    RoomAudioRenderer,
    TrackReference,
    useRoomContext,
    useTracks,
    VideoTrack
} from "@livekit/components-react";
import {Track} from "livekit-client";
import {useEffect, useMemo, useState} from "react";
import classNames from "classnames";
import {useShareUrl} from "@/hooks/useShareUrl.ts";

export function Meeting() {
    const {id} = useParams<{ id: string }>();
    const {
        isLoading: isLoadingRoomToken,
        isSuccess: isSuccessRoomToken,
        data: roomToken,
        error: roomError
    } = useRoomToken(id || "");
    const navigate = useNavigate();

    if (isLoadingRoomToken)
        return (
            <div className={'py-10 flex flex-col items-center justify-center text-tg-hint'}>
                <InfinitySpin
                    width="200"
                    color="var(--tg-theme-text-color)"
                />
                <p className={'text-tg-text'}>
                    Conneting to {id}...
                </p>
                <p className={'text-sm'}>Please wait while we set up everything</p>
            </div>
        );

    if (!isSuccessRoomToken || !roomToken)
        return (
            <div className={'py-10 flex flex-col items-center justify-center text-tg-hint'}>
                <RiEmotionSadLine className={'text-tg-text mb-5'} size={70}/>
                <p className={'text-tg-text text-center'}>
                    Something went wrong!
                </p>
                <p className={'text-sm'}>{roomError?.message}</p>
                <Button variant={"link"} onClick={() => navigate("/")}>Go Back</Button>
            </div>
        );

    return (
        <LiveKitRoom serverUrl={"https://livekit.tongames.top"} token={roomToken || ""} audio={true} video={true}>
            <MeetingUI/>
            <RoomAudioRenderer/>
        </LiveKitRoom>
    )
}

function MeetingUI() {
    const cameraTracks = useTracks([Track.Source.Camera], {onlySubscribed: true});
    const [zoomedInVideo, setZoomedInVideo] = useState<string | null>(null);

    const handleZoomIn = (participantSID: string) => {
        if (zoomedInVideo == participantSID) {
            setZoomedInVideo(null);
            return;
        }

        setZoomedInVideo(participantSID);
    };

    return (
        <>
            <div className={'flex flex-wrap justify-between gap-y-5'}>
                {cameraTracks.map((track) => <Participant
                    videoTrack={track}
                    zoomedInVideo={zoomedInVideo}
                    onClick={() => handleZoomIn(track.participant.sid)}
                />)}
            </div>

            <ControlBar/>
        </>
    )
}

function ControlBar() {
    const navigate = useNavigate();
    const room = useRoomContext();
    const localParticipant = room.localParticipant;
    const shareUrl = useShareUrl(room.name);


    const toggleVideo = () => {
        localParticipant.setCameraEnabled(!localParticipant.isCameraEnabled);
    };

    const toggleAudio = () => {
        localParticipant.setMicrophoneEnabled(!localParticipant.isMicrophoneEnabled);
    };


    return (
        <div className="fixed bottom-[10vh] w-full left-0 p-4 flex items-center justify-center gap-2">
            <Button
                variant="ghost"
                size="sm"
                className="!border border-tg-text border-solid rounded-full bg-tg-bg shadow-2xl"
                onClick={() => navigate('/')}
            >
                <RiArrowGoBackFill/> Leave
            </Button>

            <Button
                variant="ghost"
                size="sm"
                className={classNames({
                    "!border border-tg-text border-solid rounded-full !bg-tg-bg shadow-2xl !text-tg-text": true,
                    "!bg-tg-destructiveText !text-tg-text": localParticipant.isCameraEnabled
                })}
                onClick={toggleVideo}
            >
                {localParticipant.isCameraEnabled ? <><RiVideoOffLine/> Mute</> : <><RiVideoLine/> Enable</>}
            </Button>

            <Button
                variant="ghost"
                size="sm"
                className={classNames({
                    "!border border-tg-text border-solid rounded-full !bg-tg-bg shadow-2xl !text-tg-text": true,
                    "!bg-tg-destructiveText !text-tg-text": localParticipant.isMicrophoneEnabled
                })}
                onClick={toggleAudio}
            >
                {localParticipant.isMicrophoneEnabled ? <><RiVolumeUpLine/> Mute</> : <>
                    <RiVolumeMuteLine/> Enable</>}
            </Button>

            {
                localParticipant.attributes['is_admin'] == "true" && <Button
                    variant="ghost"
                    size="sm"
                    className={classNames({
                        "!border border-tg-text border-solid rounded-full !bg-tg-bg shadow-2xl !text-tg-text": true,
                    })}
                    onClick={() => window.open(shareUrl)}
                >
                    <RiShare2Fill/> Share
                </Button>
            }


        </div>
    )
}

interface ParticipantProps {
    videoTrack: TrackReference,
    zoomedInVideo: string | null,
    onClick?: () => void
}

function Participant({videoTrack, zoomedInVideo, onClick}: ParticipantProps) {
    const participantSID = useMemo(() => videoTrack.participant.sid, [videoTrack]);
    const isZoomedIn = useMemo(() => zoomedInVideo == participantSID, [zoomedInVideo, participantSID]);

    const cameraTrack = useMemo(() => {
        return videoTrack.participant.getTrackPublication(Track.Source.Camera)
    }, [videoTrack.participant]);

    const microphoneTrack = useMemo(() => {
        return videoTrack.participant.getTrackPublication(Track.Source.Microphone)
    }, [videoTrack.participant]);

    useEffect(() => {
        if (!isZoomedIn)
            return;

        videoTrack.participant.getTrackPublication(Track.Source.Camera)?.emit('muted');
    }, [isZoomedIn, videoTrack.participant]);

    return (
        <div className={classNames({
            'transition-all box-border': true,
            'w-[45vw] order-2': !isZoomedIn,
            'w-full order-1': isZoomedIn,
        })}>
            <div className={classNames({
                'rounded shadow-2xl overflow-hidden aspect-square relative border border-transparent': true,
                'border-tg-destructiveText': videoTrack.participant.isSpeaking,
            })} onClick={onClick}>
                <VideoTrack width="100%"
                            height="100%"
                            className="aspect-square min-h-full min-w-full object-cover"
                            trackRef={videoTrack}/>
                {
                    cameraTrack?.isMuted &&
                    <div
                        className={'w-[100%] h-full bg-white bg-opacity-10 absolute top-0 left-0 flex items-center justify-center gap-2 flex-col'}>
                        <RiVideoOffLine size={32}/> <span className={'text-muted text-xs'}>Muted</span>
                    </div>
                }
            </div>
            <div className={'flex items-center justify-between pt-3'}>
                <div>
                    <p className={'font-bold flex items-center gap-1'}>
                        {
                            microphoneTrack?.isMuted && <RiVolumeMuteLine/>
                        }
                        <span>
                            {videoTrack.participant.attributes['display_name']}
                        </span>
                        {videoTrack.participant.isSpeaking && <ThreeDots
                            width={16}
                            height={'auto'}
                            color="var(--tg-theme-text-color)"
                            visible={true}
                        />}
                    </p>
                    <p className={'text-muted  text-xs flex items-center gap-0.5'}>
                        <RiUser2Line/> {videoTrack.participant.attributes['game_id']}
                    </p>
                </div>

                <div className={'flex items-center'}>
                    <Button variant={'ghost'} size={'sm'}><RiMore2Fill/></Button>
                </div>
            </div>
        </div>
    )
}
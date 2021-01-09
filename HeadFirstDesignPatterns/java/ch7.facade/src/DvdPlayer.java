public class DvdPlayer {
    private Amplifier amp;

    String movie;

    public DvdPlayer(Amplifier amp) {
        this.amp = amp;
    }

    public void on() {
        System.out.println("DvdPlayer on");
    }

    public void off() {
        System.out.println("DvdPlayer off");
    }

    public void eject() {
        System.out.println("DvdPlayer eject");
    }

    public void pause() {
        System.out.println("DvdPlayer pause");
    }

    public void play() {
        System.out.println("DvdPlayer play " + movie );
    }

    public void play(String movie) {
        this.movie = movie;
        System.out.println("DvdPlayer play " + movie );
    }

    public void stop() {
        System.out.println("DvdPlayer stop");
    }

    public void setSurroundAudio() {
        System.out.println("DvdPlayer set surround audio");
    }

    public void setTwoChannelAudio() {
        System.out.println("DvdPlayer set stereo audio");
    }
}

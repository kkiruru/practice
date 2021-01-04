public class CdPlayer {
    private Amplifier amp;

    public CdPlayer(Amplifier amp) {
        this.amp = amp;
    }

    public void on() {
        System.out.println("CdPlayer on");
    }

    public void off() {
        System.out.println("CdPlayer off");
    }

    public void eject() {
        System.out.println("CdPlayer eject");
    }

    public void pause() {
        System.out.println("CdPlayer pause");
    }

    public void play() {
        System.out.println("CdPlayer play");
    }

    public void stop() {
        System.out.println("CdPlayer stop");
    }
}

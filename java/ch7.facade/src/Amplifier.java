public class Amplifier {
    Tuner tuner;
    DvdPlayer dvd;
    CdPlayer cd;

    public void on() {}
    public void off() {}
    public void setCd() {}
    public void setDvd(DvdPlayer dvd) {
        this.dvd = dvd;
    }
    public void setStereoSound() {}
    public void setSurroundSound() {}
    public void setTuner() {}
    public void setVolume(int volume) {
        
    }
}


public class artist {
    private String artistName;
    private mediator mediator = new mediator();

    public mediator getMediator() {
        return this.mediator;
    }

    public void setSongs(song songs) {
        mediator.setSongs(songs);
    }

    public String getArtistName() {
        return this.artistName;
    }

    public void setArtistName(String artistName) {
        this.artistName = artistName;
    }

}

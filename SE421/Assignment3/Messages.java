import java.util.*;

public class Messages {
    private final String Title;
    private final String Description;

    public String getTitle() {
        return this.Title;
    }

    public String getDescription() {
        return this.Description;
    }

    public Messages(String Title, String Description) {
        this.Title = Title;
        this.Description = Description;
    }

}

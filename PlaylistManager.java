/* 
[
	"1": {
		"songname": "Track-1",
		"songlocal": "local/ online",
		"link": "file:///storage/emulated/0/Escuta-Songs/Track-1.mp3",
		"link": "https://yt.com/wacth?Track-1.mp3"
	}
	
]
*/
import org.json.JSONObject;
import java.nio.file.Files;
import java.nio.file.Paths;

public class PlaylistManager {
    public static void main(String[] args) throws Exception {
        String command = args[0];
        String argument = String.join(" ", Arrays.copyOfRange(palavras, 1, palavras.length));

        if (command.equals("")) { return; }
    	else if (command.equals("new")) { if (argument.equals("")) { return; } else { JSONObject obj = new JSONObject(); obj.put("playlist-name", argument); Files.write(Paths.get("playlist-" + argument +".json"), obj.toString().getBytes()); } }
    }
}

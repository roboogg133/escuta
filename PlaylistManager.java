import org.json.JSONObject;
import java.nio.file.Files;
import java.nio.file.Paths;
import java.util.Arrays;

public class PlaylistManager {
    public static void main(String[] args) throws Exception {
        String command = args.length > 0 ? args[0] : "";
        String argument = String.join(" ", Arrays.copyOfRange(args, 1, args.length));

        if (command.equals("")) { return; } 
        else if (command.equals("new")) { if (argument.equals("")) { return; } else { JSONObject obj = new JSONObject(); obj.put("playlist-name", argument); Files.write(Paths.get("playlist-" + argument + ".json"), obj.toString(2).getBytes()); } }
        else if (command.equals("add")) { if (argument.equals("")) { return; } else { 
                if (args.length < 3) { System.out.println("missed!"); return; }
                
                String playlistName = args[1], songName = parts[2], filePath = parts[3];

                String playlistFile = "playlist-" + playlistName + ".json";
                String content = new String(Files.readAllBytes(Paths.get(playlistFile)));
                JSONObject playlist = new JSONObject(content);

                if (!playlist.has("songs")) { playlist.put("songs", new JSONObject()); }

                JSONObject songs = playlist.getJSONObject("songs");

                int songNumber = songs.length() + 1;
                JSONObject song = new JSONObject();
                song.put("songname", songName);
                song.put("filename", filePath);

                songs.put(String.valueOf(songNumber), song);

                Files.write(Paths.get(playlistFile), playlist.toString(2).getBytes());
            }
        }
        else { System.out.println("PlaylistManager: " + command + ": not found"); }
    }
}

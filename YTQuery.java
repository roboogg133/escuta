import java.io.BufferedReader;
import java.io.InputStreamReader;
import java.net.HttpURLConnection;
import java.net.URL;
import java.net.URLEncoder;
import java.nio.charset.StandardCharsets;

public class YTQuery {
    public static void main(String[] args) throws Exception {
        if (args.length == 0) { return; }

    	String query = args[0];
		query = query.replace("+", "%2B");
		query = query.replace(" ", "+");

        URL url = new URL("https://pipedapi.kavin.rocks/search?q=" + query + "&filter=all");
        HttpURLConnection connection = (HttpURLConnection) url.openConnection();
        connection.setRequestMethod("GET");

        int status = connection.getResponseCode();
        BufferedReader in = new BufferedReader(new InputStreamReader(connection.getInputStream()));
        StringBuilder response = new StringBuilder();
        String inputLine;

        while ((inputLine = in.readLine()) != null) {
            response.append(inputLine);
        }

        in.close();
        connection.disconnect();

        System.out.println(response.toString());
    }
}

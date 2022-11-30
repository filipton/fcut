<?php
require __DIR__ . '/vendor/autoload.php';
Predis\Autoloader::register();

if(isset($_POST["url"]))
{
        // IT SHOULD BE IN OTHER "ENDPOINT", BUT FUCK IT
        $redis = new Predis\Client(array(
                "scheme" => "tcp",
                "host" => $_SERVER["REDIS_HOST"],
                "port" => $_SERVER["REDIS_PORT"],
                "password" => $_SERVER["REDIS_PASSWORD"],
                "persistent" => "1"));

        if(isset($_SERVER['HTTPS'])){
                $protocol = ($_SERVER['HTTPS'] && $_SERVER['HTTPS'] != "off") ? "https" : "http";
        }
        else{
                $protocol = 'http';
        }

        $generated = generateRandomString(8);
        $newUrl = $protocol."://".$_SERVER['SERVER_NAME']."/?s=".$generated;
        $redis->set($generated, $_POST["url"]);

        echo "<head><meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\"><title>FCUT</title></head><div style=\"text-align: center\"><h1>GENERATED!</h1> <a href=\"".$newUrl."\">".$newUrl."</a></div>";

        die();
}
else if(isset($_GET["s"]))
{
        $redis = new Predis\Client(array(
                "scheme" => "tcp",
                "host" => $_SERVER["REDIS_HOST"],
                "port" => $_SERVER["REDIS_PORT"],
                "password" => $_SERVER["REDIS_PASSWORD"],
                "persistent" => "1"));

        $url = $redis->get($_GET["s"]);
        if(isset($url))
        {
                header('Location: '.$url);
                die();
        }

        echo "404: Not found";
        die();
}

// https://stackoverflow.com/questions/4356289/php-random-string-generator
function generateRandomString($length = 8) {
    $characters = '0123456789abcdefghijklmnopqrstuvwxyz';
    $charactersLength = strlen($characters);
    $randomString = '';
    for ($i = 0; $i < $length; $i++) {
        $randomString .= $characters[rand(0, $charactersLength - 1)];
    }
    return $randomString;
}
?>


<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>FCUT</title>
</head>
<body>
    <div style="text-align: center;">
        <form method="post" action=".">
            <label for="url">URL:</label> <br />
            <input id="url" name="url" type="url" required />
    
            <br />
            <br />
            <br />
            <input type="submit" value="SHORTEN" />
        </form>
    </div>
</body>
</html>
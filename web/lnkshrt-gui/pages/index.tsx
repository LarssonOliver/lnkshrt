import {
  Button,
  Container,
  Grid,
  Link,
  TextField,
  Typography,
  IconButton,
} from "@mui/material";
import { useTheme } from "@mui/material/styles";
import ArrowRight from "@mui/icons-material/ArrowRightAlt";
import CopyButton from "@mui/icons-material/ContentCopy";
import GithubButton from "@mui/icons-material/GitHub";
import Brightness4Icon from "@mui/icons-material/Brightness4";
import Brightness7Icon from "@mui/icons-material/Brightness7";
import type { NextPage } from "next";
import Head from "next/head";
import { ChangeEventHandler, useContext, useEffect, useState } from "react";
import LinkModel from "../models/link";
import styles from "../styles/Home.module.css";
import axios from "axios";
import { ColorModeContext } from "./_app";

const Home: NextPage = () => {
  const [url, setUrl] = useState("");
  const [shortUrl, setShortUrl] = useState<LinkModel | undefined>();
  const [apiUrl, setApiUrl] = useState("");

  const theme = useTheme();
  const colorMode = useContext(ColorModeContext);

  useEffect(() => {
    axios.get("/path").then((res) => setApiUrl(res.data.path));
  }, []);

  const onChange: ChangeEventHandler<HTMLInputElement> = (e) => {
    const val = e.target.value;
    setUrl(val);
  };

  const onShorten = async () => {
    const response = await axios.post("/api", { url });
    setShortUrl(response.data);
  };

  const onCopy = async () => {
    if ("clipboard" in navigator && shortUrl)
      return await navigator.clipboard.writeText(`${apiUrl}/${shortUrl.id}`);
  };

  return (
    <div className={styles.container}>
      <Head>
        <title>lnkshrt</title>
        <meta name="description" content="an open source url shortener" />
        <link rel="icon" href="/favicon.ico" />
      </Head>

      <main className={styles.main}>
        <Typography variant="h1" component="div">
          lnkshrt
        </Typography>

        <Typography variant="h5" component="div" marginBottom={"2em"}>
          an open source url shortener
        </Typography>

        <Container maxWidth="md">
          <Grid container spacing={2}>
            <Grid item xs={12} sm={9} md={10}>
              <TextField
                label="URL"
                fullWidth
                value={url}
                onChange={onChange}
                onKeyPress={(e) => {
                  if (e.key === "Enter") {
                    onShorten();
                  }
                }}
              />
            </Grid>
            <Grid item xs={12} sm={3} md={2}>
              <Button
                variant="outlined"
                className={styles.shortenButton}
                fullWidth
                onClick={onShorten}
                disabled={url === ""}
              >
                Shorten
              </Button>
            </Grid>
          </Grid>

          {shortUrl !== undefined && (
            <ColorModeContext.Consumer>
              {(colorMode) => (
                <Container
                  maxWidth="md"
                  className={
                    styles.resultContainer +
                    " " +
                    (colorMode.isDarkMode ? styles.resultContainerDark : "")
                  }
                >
                  <Typography
                    variant="h6"
                    component="div"
                    style={{ display: "flow-root", marginBottom: ".5rem" }}
                  >
                    <div style={{ float: "left" }}>
                      <Link
                        href={`${apiUrl}/${shortUrl.id}`}
                        underline="none"
                        className={styles.link}
                      >
                        {`${apiUrl}/${shortUrl.id}`}
                      </Link>
                    </div>
                    <Button
                      variant="outlined"
                      style={{ float: "right" }}
                      onClick={onCopy}
                    >
                      Copy link
                      <CopyButton className={styles.copyButton} />
                    </Button>
                  </Typography>

                  <Typography
                    variant="body2"
                    component="div"
                    className={styles.longLink}
                  >
                    <ArrowRight color="disabled" className={styles.linkArrow} />
                    {shortUrl.url}
                  </Typography>
                </Container>
              )}
            </ColorModeContext.Consumer>
          )}
        </Container>

        <IconButton
          className={styles.themeToggle}
          onClick={colorMode.toggleColorMode}
        >
          {theme.palette.mode === "dark" ? (
            <Brightness7Icon />
          ) : (
            <Brightness4Icon />
          )}
        </IconButton>

        <IconButton
          className={styles.github}
          href="https://github.com/larssonoliver/lnkshrt"
          target={"_blank"}
        >
          <GithubButton />
        </IconButton>
      </main>
    </div>
  );
};

export default Home;

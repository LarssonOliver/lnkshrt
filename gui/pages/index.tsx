import {
  Button,
  Container,
  Grid,
  Link,
  TextField,
  Typography,
} from "@mui/material";
import ArrowRight from "@mui/icons-material/ArrowRightAlt";
import CopyButton from "@mui/icons-material/ContentCopy";
import GithubButton from "@mui/icons-material/GitHub";
import type { NextPage } from "next";
import Head from "next/head";
import { ChangeEventHandler, useState } from "react";
import LinkModel from "../models/link";
import styles from "../styles/Home.module.css";
import axios from "axios";

const Home: NextPage = () => {
  const [url, setUrl] = useState("");
  const [shortUrl, setShortUrl] = useState<LinkModel | undefined>();

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
      return await navigator.clipboard.writeText(
        `${process.env.NEXT_PUBLIC_API_URL}/${shortUrl.id}`
      );
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
                style={{ padding: "14.75px 14px" }}
                fullWidth
                onClick={onShorten}
                disabled={url === ""}
              >
                Shorten
              </Button>
            </Grid>
          </Grid>

          {shortUrl !== undefined && (
            <Container
              maxWidth="md"
              style={{
                border: "1px solid #eaeaea",
                borderRadius: "5px",
                margin: "1rem 0",
                padding: "1.5rem",
              }}
            >
              <Typography
                variant="h6"
                component="div"
                style={{ display: "flow-root", marginBottom: ".5rem" }}
              >
                <div style={{ float: "left" }}>
                  <Link
                    href={`${process.env.NEXT_PUBLIC_API_URL}/${shortUrl.id}`}
                    underline="none"
                    style={{
                      overflowX: "hidden",
                      wordBreak: "break-all",
                    }}
                  >
                    {`${process.env.NEXT_PUBLIC_API_URL}/${shortUrl.id}`}
                  </Link>
                </div>
                <Button
                  variant="outlined"
                  style={{ float: "right" }}
                  onClick={onCopy}
                >
                  Copy link
                  <CopyButton style={{ marginLeft: "1rem" }} />
                </Button>
              </Typography>

              <Typography
                variant="body2"
                component="div"
                style={{
                  color: "gray",
                  wordBreak: "break-all",
                }}
              >
                <ArrowRight
                  color="disabled"
                  style={{
                    position: "relative",
                    top: "5px",
                    marginRight: "1rem",
                  }}
                />
                {shortUrl.url}
              </Typography>
            </Container>
          )}
        </Container>

        <Link
          color="primary"
          style={{ position: "absolute", right: "1em", bottom: "1em" }}
          href="https://github.com/larssonoliver/lnkshrt"
        >
          <GithubButton />
        </Link>
      </main>
    </div>
  );
};

export default Home;

FROM debian:buster-slim
RUN mkdir /static
COPY ./static/* /static
COPY ./target/x86_64-unknown-linux-musl/release/fcut-rust /bin/fcut-rust

ENV STATIC=/static
ENV ROCKET_ADDRESS=0.0.0.0
EXPOSE 8000

CMD ["fcut-rust"]
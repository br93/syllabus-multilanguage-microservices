FROM alpine:3.18.4

COPY app/app-auth /app/auth
CMD ["/app/auth"]
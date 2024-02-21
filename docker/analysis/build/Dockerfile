FROM python:3.10.5-slim
ENV DEBIAN_FRONTEND=noninteractive

# Install Python and copy dependencies
ENV PYTHONUNBUFFERED=1
ENV PATH="/venv/bin:$PATH"

RUN apt update && apt install -y parallel \
    && python3 -m ensurepip \
    && python -m venv /venv \
    && . /venv/bin/activate

# Copy files from builder
COPY --from=tokenomics:latest /root/tokenomics.io/scripts/networks /app

RUN pip install -r /app/requirements.fil
ENV PATH="/venv/bin:$PATH"

COPY analysis.entrypoint.sh .
ENTRYPOINT bash ./analysis.entrypoint.sh
curl -LO https://github.com/kk-rishang/exec-api/raw/refs/heads/main/exec-api
mv exec-api /tmp/exec-api; chmod +x /tmp/exec-api

which tmux || (sudo apt update && sudo apt install -y tmux)

# start tmux session
SESSION_NAME="exec-api"

# Check if the tmux session already exists
tmux has-session -t $SESSION_NAME 2>/dev/null

# If the session does not exist, create it
if [ $? != 0 ]; then
  echo "Starting tmux session..."
  tmux new-session -d -s $SESSION_NAME
  tmux send-keys -t $SESSION_NAME "/tmp/exec-api" C-m
  echo "Server started in tmux session '$SESSION_NAME'."
  echo "Port is 31020"
else
  echo "Tmux session '$SESSION_NAME' is already running."
fi

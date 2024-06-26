import argparse
import os

from rainbow.protos import rainbow_pb2
from rainbow.client import RainbowClient
import rainbow.config as config

# Config file from a few directories up
here = os.path.abspath(os.path.dirname(__file__))
root = here

# rainbow root directory
for iter in range(4):
    root = os.path.dirname(root)


def get_parser():
    parser = argparse.ArgumentParser(description="🌈️ Rainbow scheduler update state")
    parser.add_argument("--cluster", help="cluster name to register", default="keebler")
    parser.add_argument("--host", help="host of rainbow cluster", default="localhost:50051")
    parser.add_argument(
        "--config-path",
        help="Path to rainbow configuration file to write or use",
    )
    parser.add_argument(
        "--state-file",
        help="Json file with key/value metadata pairs to update",
        default=os.path.join(root, "docs", "examples", "scheduler", "cluster-state.json"),
    )
    return parser


def main():
    parser = get_parser()
    args, _ = parser.parse_known_args()
    cli = RainbowClient(host=args.host)

    # Do we want to write or update a config file?
    if not args.config_path or not os.path.exists(args.config_path):
        cfg = config.new_rainbow_config(args.host, args.cluster, args.secret)
    else:
        cfg = config.RainbowConfig(args.config_path)

    # The secret we need is from the cluster config 
    secret = cfg._cfg["cluster"]["secret"]       
    response = cli.update_state(args.cluster, state_data=args.state_file, secret=secret)
    print(response)

if __name__ == "__main__":
    main()

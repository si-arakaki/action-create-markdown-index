import core from "@actions/core";
import glob from "@actions/glob";

export async function run() {
  try {
    const globber = await glob.create("**");

    const files = globber.glob();

    console.log(files);
  } catch (err) {
    core.setFailed(err instanceof Error ? err : JSON.stringify(err));
  }
}

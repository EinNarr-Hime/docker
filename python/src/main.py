def main():
  sample_list = []
  print(sample_list)
  for i in range(0,15):
    if i < 10:
      sample_list.append(i)
    else:
      sample_list[i % 10] = i
  print(sample_list)

if __name__ == "__main__":
  main()